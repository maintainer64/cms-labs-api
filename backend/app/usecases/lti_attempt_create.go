package usecases

import (
	"errors"
	"fmt"
	"net/url"
	"time"

	"gitlab.com/a10869/api-modules/shared/jsonrpc"

	"gitlab.com/a10869/api-modules/shared/cms_client"

	json "github.com/goccy/go-json"
	"github.com/ory/go-convenience/mapx"
	"github.com/rs/zerolog/log"
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/queries"
	"gitlab.com/a10869/api-modules/backend/app/queries/lti_query"
)

type LTIAttemptCreateUC struct {
	LTIAttemptQueries     *queries.LTIAttemptQueries
	LTIRoomQueries        *queries.LTIRoomQueries
	LaunchData            lti_query.LaunchDataStorer
	RoundQueuePoolQueries *queries.RoundQueuePoolQueries
	LTIRoutingQueries     *queries.LTIRoutingQueries
	PNETServerQueries     *queries.PNETServerQueries
	UserQueries           *queries.UserQueries
	user                  *cms_client.SSOTokenPublicData
}

type LTIAttemptCreateInputDTO struct {
	RoomNumber int64 `json:"room_number"`
}

type LTIAttemptCreateRequest struct {
	JSONRPC string                   `json:"jsonrpc" default:"2.0" validate:"required"`
	Method  string                   `json:"method" default:"lti_attempt.create" validate:"required"`
	Params  LTIAttemptCreateInputDTO `json:"params,omitempty"`
	ID      string                   `json:"id,omitempty" default:"1" validate:"required"`
}

type LTIAttemptCreateOutputDTO struct {
	RoomNumber    int64                 `json:"room_number"`
	Collaboration int                   `json:"collaboration"`
	NextUrl       string                `json:"next_url"`
	AutoRedirect  bool                  `json:"auto_redirect"`
	Members       []models.UserListItem `json:"members"`
}

type LTIAttemptCreateResponse struct {
	JSONRPC string                    `json:"jsonrpc" default:"2.0" validate:"required"`
	Result  LTIAttemptCreateOutputDTO `json:"result,omitempty"`
	Error   interface{}               `json:"error,omitempty"`
	ID      string                    `json:"id,omitempty" default:"1" validate:"required"`
}

func (u *LTIAttemptCreateUC) SetContext(user *cms_client.SSOTokenPublicData) *LTIAttemptCreateUC {
	u.user = user
	return u
}

func (u *LTIAttemptCreateUC) Execute(dto LTIAttemptCreateInputDTO) (LTIAttemptCreateOutputDTO, error) {
	if u.user == nil {
		return LTIAttemptCreateOutputDTO{}, errors.New("not logged in")
	}
	log.Info().Msg(
		fmt.Sprintf(
			"LTIAttemptCreateUC: Get or create attempt on tasks pnet-servers by user: %+v",
			u.user.Email,
		),
	)
	route := u.SearchRelevantRouting()
	if route == nil || route.ID == 0 {
		return LTIAttemptCreateOutputDTO{}, queries.LTIRoutingNotFoundError
	}
	if route.PNETLabsType == cms_client.PNETLabsTypeSSO {
		return LTIAttemptCreateOutputDTO{
			AutoRedirect: true,
		}, nil
	}
	attempt, _ := u.LTIAttemptQueries.GetActiveByUserId(u.user.UserID(), route.ID)

	if attempt.ID != 0 && attempt.RoomID != nil && dto.RoomNumber != 0 {
		log.Info().Msg(
			fmt.Sprintf(
				"LTIAttemptCreateUC: Change room from roomID: %d to %d by id %s",
				attempt.RoomID,
				dto.RoomNumber,
				attempt.AttemptID,
			),
		)
		roomEntity, err := u.LTIRoomQueries.GetByRoomNumber(dto.RoomNumber)
		if err != nil {
			return LTIAttemptCreateOutputDTO{}, jsonrpc.NewRpcError("not_found_room", "not found other room")
		}
		otherAttempts, _ := u.LTIAttemptQueries.GetByRoomID(roomEntity.ID)
		if len(otherAttempts) == 0 {
			return LTIAttemptCreateOutputDTO{}, jsonrpc.NewRpcError("not_found_room", "not change room")
		}
		otherAttempt := otherAttempts[0]
		attempt.PNETServerID = otherAttempt.PNETServerID
		attempt.RoomID = otherAttempt.RoomID
		attempt.LTIRoutingID = otherAttempt.LTIRoutingID
		attempt.ExtendExpiredAt(1)
		if err := u.LTIAttemptQueries.Upsert(&attempt); err != nil {
			return LTIAttemptCreateOutputDTO{}, err
		}
		return u.PreparedResponseByAttempt(&attempt)
	}
	if attempt.ID != 0 {
		log.Info().Msg(
			fmt.Sprintf(
				"LTIAttemptCreateUC: Getted last not expired attempt on tasks pnet-servers by user: %+v",
				u.user.Email,
			),
		)
		attempt.ExtendExpiredAt(1)
		if err := u.LTIAttemptQueries.Upsert(&attempt); err != nil {
			return LTIAttemptCreateOutputDTO{}, err
		}
		return u.PreparedResponseByAttempt(&attempt)
	}
	attempt = models.LTIAttempt{}
	attempt.UserID = u.user.UserID()
	// Set LTIRoutingSecretID
	attempt.LTIRoutingID = route.ID
	// Set RoomID
	if route.Collaboration > 1 {
		roomEntity, err := u.LTIRoomQueries.Create()
		if err != nil || roomEntity == nil {
			log.Warn().Msg(
				fmt.Sprintf(
					"LTIAttemptCreateUC: Not created room. Error: %+v",
					err,
				),
			)
			return LTIAttemptCreateOutputDTO{}, jsonrpc.NewRpcError("not_found_room", "not created room")
		}
		attempt.RoomID = &roomEntity.ID
	}
	// Set ExpiredAt
	if route.PinnedSessionMinutes > 1 {
		attempt.ExpiredAt = time.Now().UTC().Add(time.Duration(route.PinnedSessionMinutes) * time.Minute)
	} else {
		attempt.ExpiredAt = time.Now().UTC().Add(time.Duration(1) * time.Hour)
	}
	if err := u.LTIAttemptQueries.Upsert(&attempt); err != nil {
		return LTIAttemptCreateOutputDTO{}, err
	}
	return u.PreparedResponseByAttempt(&attempt)
}

// AllocatedServer - Функция выделения сервера по правилу из Route или по общему пулу серверов для всей
// комнаты подключенной к данной попытки
func (u *LTIAttemptCreateUC) AllocatedServer(attempt *models.LTIAttempt) (models.PNETServer, error) {
	// Получаем блокировку
	unlockFn, err := u.LTIAttemptQueries.AllocatedServerLock(attempt.ID, attempt.RoomID)
	if err != nil {
		return models.PNETServer{}, err
	}
	defer unlockFn()

	// Закрепленная сессия
	if attempt.PNETServerID != nil && *attempt.PNETServerID != 0 {
		pnetServer, err := u.PNETServerQueries.Get(*attempt.PNETServerID)
		if err == nil {
			log.Info().Msg(
				fmt.Sprintf("attemptID %d with pnet_server_id %d already allocation",
					attempt.ID,
					*attempt.PNETServerID,
				),
			)
			return pnetServer, nil
		}
		log.Info().Msg(
			fmt.Sprintf("attemptID %d with pnet_server_id %d not found with allocation",
				attempt.ID,
				*attempt.PNETServerID,
			),
		)
	}
	// Правило маршрута
	if attempt.LTIRoutingID != 0 {
		ltiRoute, err := u.LTIRoutingQueries.Get(attempt.LTIRoutingID)
		if err != nil {
			return models.PNETServer{}, err
		}
		// Правило маршрута установлено
		if ltiRoute.PNETServerID != 0 {
			pnetServer, err := u.PNETServerQueries.Get(ltiRoute.PNETServerID)
			if err != nil {
				return models.PNETServer{}, err
			}
			log.Info().Msg(
				fmt.Sprintf("attemptID %d with pnet_server_id %d allocation by route",
					attempt.ID,
					pnetServer.ID,
				),
			)
			// Закрепляем сервер
			if err := u.LTIAttemptQueries.AllocatedServer(
				attempt.ID,
				attempt.RoomID,
				pnetServer.ID,
			); err != nil {
				return models.PNETServer{}, err
			}
			return pnetServer, nil
		}
	}
	// Выделение из пула
	pnetRoutes, err := u.RoundQueuePoolQueries.GetNextByType(models.RoundQueuePoolTypePNET)
	if err != nil {
		return models.PNETServer{}, err
	}
	pnetServer, err := u.PNETServerQueries.Get(pnetRoutes.PNETServerID)
	if err != nil {
		return models.PNETServer{}, err
	}
	log.Info().Msg(
		fmt.Sprintf("attemptID %d with pnet_server_id %d allocated from pool",
			attempt.ID,
			pnetServer.ID,
		),
	)
	// Закрепляем сервер
	if err := u.LTIAttemptQueries.AllocatedServer(
		attempt.ID,
		attempt.RoomID,
		pnetServer.ID,
	); err != nil {
		return models.PNETServer{}, err
	}
	return pnetServer, nil
}

func (u *LTIAttemptCreateUC) SearchRelevantRouting() *models.LTIRouting {
	if u.user == nil || u.user.LastLaunchId == "" {
		return nil
	}
	log.Info().Msg(fmt.Sprintf("LTIAttemptCreateUC: SearchRelevantRouting by user: %+v", u.user.Email))
	var jwtTokenPayload map[string]interface{}
	launchData, err := u.LaunchData.FindLaunchData(u.user.LastLaunchId)
	if err != nil {
		log.Info().Msg(fmt.Sprintf("LTIAttemptCreateUC: SearchRelevantRouting by user: %+v not found in db", u.user.Email))
		return nil
	}
	if err := json.Unmarshal(launchData, &jwtTokenPayload); err != nil {
		log.Info().Msg(fmt.Sprintf("LTIAttemptCreateUC: SearchRelevantRouting by user: %+v not parse json", u.user.Email))
		return nil
	}
	log.Debug().Msg(fmt.Sprintf("LTIAttemptCreateUC: SearchRelevantRouting launchData: %+v", jwtTokenPayload))
	ltiTaskId := ""
	ltiCourseId := ""
	ltiSubId := ""
	title := ""
	description := ""
	customParams := make([]string, 0)
	if resourceLinkMap, ok := jwtTokenPayload["https://purl.imsglobal.org/spec/lti/claim/resource_link"].(map[string]interface{}); ok {
		resourceLink := make(map[interface{}]interface{})
		for key, value := range resourceLinkMap {
			resourceLink[key] = value
		}
		ltiTaskId = mapx.GetStringDefault(resourceLink, "id", "")
		title = mapx.GetStringDefault(resourceLink, "title", "")
		description = mapx.GetStringDefault(resourceLink, "description", "")
	}
	if customParamsPrune, ok := jwtTokenPayload["https://purl.imsglobal.org/spec/lti/claim/custom"].(map[string]interface{}); ok {
		for k, v := range customParamsPrune {
			customParams = append(customParams, fmt.Sprintf("%s=%s", k, v))
		}
	}
	if courseLinkMap, ok := jwtTokenPayload["https://purl.imsglobal.org/spec/lti/claim/context"].(map[string]interface{}); ok {
		courseLink := make(map[interface{}]interface{})
		for key, value := range courseLinkMap {
			courseLink[key] = value
		}
		ltiCourseId = mapx.GetStringDefault(courseLink, "id", "")
	}
	if ltiSub, ok := jwtTokenPayload["sub"].(string); ok {
		ltiSubId = ltiSub
	}
	ltiRouting, err := u.LTIRoutingQueries.GetRelevantRouting(title, description, customParams)
	if err != nil || ltiRouting.ID == 0 {
		log.Warn().Msg(fmt.Sprintf("LTIAttemptCreateUC: SearchRelevantRouting exception %+v", err))
		return nil
	}
	// Обновление route добавление lti параметров курса
	if ltiRouting.LTITaskID != ltiTaskId || ltiRouting.LTICourseID != ltiCourseId || ltiRouting.LTISubID != ltiSubId {
		ltiRouting.LTICourseID = ltiCourseId
		ltiRouting.LTISubID = ltiSubId
		ltiRouting.LTITaskID = ltiTaskId
		_ = u.LTIRoutingQueries.Upsert(&ltiRouting)
	}
	return &ltiRouting
}

func (u *LTIAttemptCreateUC) PreparedResponseByAttempt(attempt *models.LTIAttempt) (LTIAttemptCreateOutputDTO, error) {
	var attempts []models.LTIAttempt
	var roomEntity models.LTIRoom
	if attempt.RoomID != nil && *attempt.RoomID > 0 {
		attempts, _ = u.LTIAttemptQueries.GetByRoomID(*attempt.RoomID)
		roomEntity, _ = u.LTIRoomQueries.Get(*attempt.RoomID)
	}
	userIds := make([]uint, 0)
	for _, item := range attempts {
		userIds = append(userIds, item.UserID)
	}
	users := make([]models.UserListItem, 0)
	if len(userIds) != 0 {
		users, _, _ = u.UserQueries.List("", userIds, len(userIds), 0)
	}
	ltiRoute, err := u.LTIRoutingQueries.Get(attempt.LTIRoutingID)
	if err != nil {
		return LTIAttemptCreateOutputDTO{}, err
	}
	waitCompleted := ltiRoute.Collaboration <= 1 || ltiRoute.Collaboration > 1 && ltiRoute.Collaboration == len(users)
	if !waitCompleted {
		return LTIAttemptCreateOutputDTO{
			RoomNumber:    roomEntity.RoomNumber,
			Collaboration: ltiRoute.Collaboration,
			NextUrl:       "#",
			AutoRedirect:  false,
			Members:       users,
		}, nil
	}
	pnetServer, err := u.AllocatedServer(attempt)
	if err != nil {
		return LTIAttemptCreateOutputDTO{}, err
	}
	nextUrl, err := u.SSOUrlGenerator(
		pnetServer.Url,
		cms_client.SSOTokenPublicExtraParams{
			AttemptID:    attempt.AttemptID,
			PNETLabsType: ltiRoute.PNETLabsType,
			PNETLabsPath: ltiRoute.PNETLabsPath,
			PNETTestPath: ltiRoute.PNETTestPath,
		},
	)
	if err != nil {
		return LTIAttemptCreateOutputDTO{}, err
	}
	return LTIAttemptCreateOutputDTO{
		RoomNumber:    roomEntity.RoomNumber,
		Collaboration: ltiRoute.Collaboration,
		NextUrl:       nextUrl,
		AutoRedirect:  true,
		Members:       users,
	}, nil
}

func (u *LTIAttemptCreateUC) SSOUrlGenerator(
	baseUrl string,
	extra cms_client.SSOTokenPublicExtraParams,
) (string, error) {
	if extra.PNETLabsType == cms_client.PNETLabsTypeClabgate {
		params := url.Values{}
		params.Add("taskId", extra.PNETLabsPath)
		encodedParams := params.Encode()
		return "/topology?" + encodedParams, nil
	}
	pathUrl, err := url.JoinPath(
		baseUrl,
		"/pnet-lab-addon/api/v1/sso/login",
	)
	nextUrl := pathUrl + "?extra=" + extra.Marshal()
	return nextUrl, err
}
