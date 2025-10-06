package external

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"

	"github.com/rs/zerolog/log"
	"gitlab.com/a10869/api-modules/backend/app/queries"
)

type PNETServerPingUC struct {
	PNETServerQueries *queries.PNETServerQueries
	LTIAttemptQueries *queries.LTIAttemptQueries
	xServiceId        string
}

type AttemptDTO struct {
	AttemptID string `json:"attempt_id" validate:"required"`
	UserEmail string `json:"user_email"`
	UserID    uint   `json:"user_id"`
}

type PNETServerPingInputDTO struct {
	Attempts []AttemptDTO `json:"attempts"`
}

type PNETServerPingRequest struct {
	JSONRPC string                 `json:"jsonrpc" default:"2.0" required:"true"`
	Method  string                 `json:"method" default:"server.ping" required:"true"`
	Params  PNETServerPingInputDTO `json:"params,omitempty"`
	ID      string                 `json:"id,omitempty" default:"1" required:"true"`
}

type PNETServerPingOutputDTO struct {
	Count int `json:"count"`
}

type PNETServerPingResponse struct {
	JSONRPC string                  `json:"jsonrpc" default:"2.0" required:"true"`
	Result  PNETServerPingOutputDTO `json:"result,omitempty"`
	Error   interface{}             `json:"error,omitempty"`
	ID      string                  `json:"id,omitempty" default:"1" required:"true"`
}

func (u *PNETServerPingUC) SetContext(xServiceId string) *PNETServerPingUC {
	u.xServiceId = xServiceId
	return u
}

func (u *PNETServerPingUC) Execute(dto PNETServerPingInputDTO) (PNETServerPingOutputDTO, error) {
	if u.xServiceId == "" {
		return PNETServerPingOutputDTO{}, errors.New("not authorized service")
	}
	log.Info().Msg("PNETServerPingUC: Update last online status and count attempts/user")
	err := u.PNETServerQueries.DB.Transaction(
		func(tx *gorm.DB) error {
			serverEntity, err := u.PNETServerQueries.GetByClientId(u.xServiceId)
			if err != nil {
				return err
			}
			now := time.Now().UTC()
			serverEntity.LastOnlineStatus = &now
			serverEntity.LastCountUsers = len(dto.Attempts)
			serverEntity.UpdatedAt = now
			tx.Save(serverEntity)
			return nil
		},
	)
	if err != nil {
		return PNETServerPingOutputDTO{}, err
	}
	count := 0
	for _, attemptDTO := range dto.Attempts {
		attempt, err := u.LTIAttemptQueries.GetByAttemptID(attemptDTO.AttemptID)
		if err != nil {
			log.Warn().Msg(fmt.Sprintf("LTI attempt id %s not found", attemptDTO.AttemptID))
			continue
		}
		attempt.ExtendExpiredAt(1)
		_ = u.LTIAttemptQueries.Upsert(&attempt)
		count++
	}
	return PNETServerPingOutputDTO{
		Count: count,
	}, nil
}
