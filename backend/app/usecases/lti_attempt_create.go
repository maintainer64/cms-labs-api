package usecases

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"net/url"
	"time"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/shared/utils"

	"gitlab.com/a10869/api-modules/shared/cms_client"

	"gitlab.com/a10869/api-modules/backend/app/usecases/response"

	"github.com/goccy/go-json"
	"github.com/ory/go-convenience/mapx"
	"github.com/rs/zerolog/log"
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/queries"
	"gitlab.com/a10869/api-modules/backend/app/queries/lti_query"
)

type LTIAttemptCreateUC struct {
	LTIAttemptQueries     *queries.LTIAttemptQueries
	LaunchData            lti_query.LaunchDataStorer
	RoundQueuePoolQueries *queries.RoundQueuePoolQueries
	LTIRoutingQueries     *queries.LTIRoutingQueries
	PNETServerQueries     *queries.PNETServerQueries
	UserQueries           *queries.UserQueries
	user                  *cms_client.SSOTokenPublicData
}

type LTIAttemptCreateInputDTO struct {
}

type LTIAttemptCreateOutputDTO struct {
	RoomNumber    *int64                `json:"room_number"`
	Collaboration int                   `json:"collaboration"`
	NextUrl       string                `json:"next_url"`
	AutoRedirect  bool                  `json:"auto_redirect"`
	Members       []models.UserListItem `json:"members"`
}

type LTIAttemptCreateResponse = response.Response[LTIAttemptCreateOutputDTO]

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
		return LTIAttemptCreateOutputDTO{}, utils.FiberValidationException{
			Status:    fiber.StatusNotFound,
			Exception: errors.New("no route found for this user"),
		}
	}
	attempt, _ := u.LTIAttemptQueries.GetActiveByUserId(u.user.Id, route.ID)
	if attempt.ID != 0 {
		log.Info().Msg(
			fmt.Sprintf(
				"LTIAttemptCreateUC: Getted last not expired attempt on tasks pnet-servers by user: %+v",
				u.user.Email,
			),
		)
		attempt.ExtendExpiredAt(2)
		_ = u.LTIAttemptQueries.Upsert(&attempt)
		return u.PreparedResponseByAttempt(attempt)
	}
	pnetRoutes, err := u.RoundQueuePoolQueries.GetNextByType(models.RoundQueuePoolTypePNET)
	if err != nil {
		return LTIAttemptCreateOutputDTO{}, err
	}
	attempt = models.LTIAttempt{}
	attempt.UserID = u.user.Id
	// Set PNETServerID
	attempt.PNETServerID = pnetRoutes.PNETServerID
	// Set LTIRoutingSecretID
	attempt.LTIRoutingID = route.ID
	// Set RoomNumber
	if route.Collaboration > 1 {
		var minRand int64 = 100000
		var maxRand int64 = 999999
		roomNumberFrom, _ := rand.Int(rand.Reader, big.NewInt(maxRand-minRand+1))
		roomNumber := roomNumberFrom.Int64() + minRand
		attempt.RoomNumber = &roomNumber
	}
	// Set ExpiredAt
	if route.PinnedSessionMinutes > 1 {
		attempt.ExpiredAt = time.Now().UTC().Add(time.Duration(route.PinnedSessionMinutes) * time.Minute)
	} else {
		attempt.ExpiredAt = time.Now().UTC().Add(time.Duration(2*60) * time.Minute)
	}
	if err = u.LTIAttemptQueries.Upsert(&attempt); err != nil {
		return LTIAttemptCreateOutputDTO{}, err
	}
	return u.PreparedResponseByAttempt(attempt)
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
	if err != nil {
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

func (u *LTIAttemptCreateUC) PreparedResponseByAttempt(attempt models.LTIAttempt) (LTIAttemptCreateOutputDTO, error) {
	attempts, err := u.LTIAttemptQueries.GetByRoomNumber(attempt.RoomNumber)
	if err != nil {
		return LTIAttemptCreateOutputDTO{}, err
	}
	userIds := make([]uint, 0)
	for _, item := range attempts {
		userIds = append(userIds, item.UserID)
	}
	users := make([]models.UserListItem, 0)
	if len(userIds) == 0 {
		users, _, _ = u.UserQueries.List("", userIds, len(userIds), 0)
	}
	pnetServer, err := u.PNETServerQueries.Get(attempt.PNETServerID)
	if err != nil {
		return LTIAttemptCreateOutputDTO{}, err
	}
	ltiRoute, err := u.LTIRoutingQueries.Get(attempt.LTIRoutingID)
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
		RoomNumber:    attempt.RoomNumber,
		Collaboration: ltiRoute.Collaboration,
		NextUrl:       nextUrl,
		AutoRedirect:  ltiRoute.Collaboration <= 1 || ltiRoute.Collaboration > 1 && ltiRoute.Collaboration == len(users),
		Members:       users,
	}, nil
}

func (u *LTIAttemptCreateUC) SSOUrlGenerator(
	baseUrl string,
	extra cms_client.SSOTokenPublicExtraParams,
) (string, error) {
	pathUrl, err := url.JoinPath(
		baseUrl,
		"/pnet-lab-addon/api/v1/sso/login",
	)
	nextUrl := pathUrl + "?extra=" + extra.Marshal()
	return nextUrl, err
}
