package usecases

import (
	"fmt"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gitlab.com/a10869/api-modules/pnetlabaddon/app/models"

	"gitlab.com/a10869/api-modules/pnetlabaddon/app/queries"
	"gitlab.com/a10869/api-modules/shared/cms_client"
)

type LabCreateInputDTO struct {
	Extra   string `json:"extra"`
	UserPod int    `json:"user_pod"`
}

type LabCreateUC struct {
	UserQueries     *queries.UserQueries
	LabSessionQuery *queries.LabSessionQuery
	*zerolog.Logger
}

func (u *LabCreateUC) CreateOrUpdateLab(
	userPod int,
	unlFilePathServer string,
	attemptId string,
) error {
	user, err := u.UserQueries.Get(userPod)
	if err != nil {
		return err
	}
	lab, _ := u.LabSessionQuery.GetByAttemptId(attemptId)
	if lab.LabSessionID != 0 {
		lab.AddJoinedUser(userPod)
		u.LabSessionQuery.Save(&lab)
		user.LabSession = &lab.LabSessionID
		u.UserQueries.Save(&user)
		log.Info().Msg(
			fmt.Sprintf(
				"UpdateLab by unlFilePathServer: %s, attemptId:%s and userPod: %d",
				unlFilePathServer,
				attemptId,
				userPod,
			),
		)
		return nil
	}
	lab = models.LabSession{
		LabSessionLID:    attemptId,
		LabSessionPod:    userPod,
		LabSessionJoined: fmt.Sprintf("%d", userPod),
		LabSessionPath:   unlFilePathServer,
	}
	u.LabSessionQuery.Save(&lab)
	user.LabSession = &lab.LabSessionID
	u.UserQueries.Save(&user)
	log.Info().Msg(
		fmt.Sprintf(
			"CreateLab by unlFilePathServer: %s, attemptId:%s and userPod: %d",
			unlFilePathServer,
			attemptId,
			userPod,
		),
	)
	return nil
}

func (u *LabCreateUC) Execute(dto LabCreateInputDTO) error {
	extra := cms_client.UnmarshalSSOTokenPublicExtraParams(dto.Extra)
	if extra.AttemptID == "" {
		u.Logger.Info().Msg("LabCreate attemptID is null")
		return nil
	}
	if extra.PNETLabsType == cms_client.PNETLabsTypeDefault && extra.PNETLabsPath != "" {
		u.Logger.Info().Msg(fmt.Sprintf("LabCreate by default path: %s", extra.PNETLabsPath))
		return u.CreateOrUpdateLab(dto.UserPod, extra.PNETLabsPath, extra.AttemptID)
	}
	return nil
}
