package lti_query

import (
	"fmt"
	"time"

	"gitlab.com/a10869/api-modules/backend/app/models"
	"gorm.io/gorm"
)

type LTINonceTokenQueries struct {
	*gorm.DB
}

func (q *LTINonceTokenQueries) ClearOldValues() error {
	timeThreshold := time.Now().UTC().Add(-5 * time.Minute)
	result := q.Where("created_at <= ?", timeThreshold).Delete(&models.LTINonceToken{})
	if result.Error != nil {
		return result.Error
	}
	log.Info().Msg(fmt.Sprintf("LTINonceTokenQueries: cleared old values %+v", timeThreshold))
	return nil
}

func (q *LTINonceTokenQueries) StoreNonce(nonce string, targetLinkURI string) error {
	if err := q.ClearOldValues(); err != nil {
		return err
	}
	entity := &models.LTINonceToken{}
	entity.ID = 0
	entity.Nonce = nonce
	entity.TargetLinkURI = targetLinkURI
	entity.CreatedAt = time.Now().UTC()
	entity.UpdatedAt = time.Now().UTC()
	log.Debug().Msg(fmt.Sprintf("LTINonceTokenQueries: store nonce %+v", entity))
	result := q.Create(entity)
	log.Info().Msg(fmt.Sprintf("LTINonceTokenQueries: stored nonce: %+v", entity.Nonce))
	return result.Error
}

func (q *LTINonceTokenQueries) TestAndClearNonce(
	nonce string,
	targetLinkURI string,
) error {
	if err := q.ClearOldValues(); err != nil {
		return err
	}
	entityDB := &models.LTINonceToken{}
	log.Debug().Msg(fmt.Sprintf(
		"LTINonceTokenQueries: check nonce: %+v and targetLinkURI: %+v",
		nonce,
		targetLinkURI,
	))
	result := q.Where("nonce = ?", nonce).Where("target_link_uri = ?", targetLinkURI).Find(&entityDB)
	log.Info().Msg(fmt.Sprintf(
		"LTINonceTokenQueries: checked nonce token: %+v",
		nonce,
	))
	if result.Error != nil {
		return result.Error
	}
	q.Where("id = ?", entityDB.ID).Delete(&entityDB)
	return nil
}
