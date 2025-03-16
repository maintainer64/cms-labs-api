package usecases

import (
	"fmt"
	"strconv"

	"github.com/rs/zerolog/log"

	"github.com/rs/zerolog"
	"gitlab.com/a10869/api-modules/pnetlabaddon/app/queries"
	"gitlab.com/a10869/api-modules/shared/cms_client"
)

type PnetServerPingUC struct {
	LabSessionQuery *queries.LabSessionQuery
	CMSClient       *cms_client.CMSClient
	*zerolog.Logger
}

func (u *PnetServerPingUC) Execute() error {
	labs, err := u.LabSessionQuery.GetRunningLabs()
	if err != nil {
		log.Warn().Msg(fmt.Sprintf("Get running labs failed: %v", err))
		return err
	}
	var attempts []cms_client.PnetServerPingAttemptDTO
	for _, lab := range labs {
		externalUserId, _ := strconv.ParseUint(lab.Name, 10, 32)
		attempt := cms_client.PnetServerPingAttemptDTO{
			AttemptID: lab.LabSessionLID,
			UserEmail: lab.Email,
			UserID:    uint(externalUserId),
		}
		attempts = append(attempts, attempt)
	}
	response, err := u.CMSClient.PnetServerPing(
		&cms_client.PNETServerPingInputDTO{
			Attempts: attempts,
		},
	)
	if response != nil {
		u.Logger.Info().Msg(fmt.Sprintf("Ping server response count: %d", response.Count))
	}
	return err
}
