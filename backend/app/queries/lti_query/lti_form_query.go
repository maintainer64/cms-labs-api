package lti_query

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/pkg/utils"
	"gorm.io/gorm"
)

type LTIFormQueries struct {
	*gorm.DB
}

func ltiGenerateKeys() (string, string, error) {
	log.Info().Msg("LTIGenerate RSA keys")
	// Generate RSA private key
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Warn().Msg(fmt.Sprintf("Error generating private key: %+v", err))
		return "", "", err
	}

	// Encode private key to PEM format
	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})

	// Generate RSA public key
	publicKey := &privateKey.PublicKey

	// Encode public key to PEM format
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		log.Warn().Msg(fmt.Sprintf("Error marshaling public key: %+v", err))
		return "", "", err
	}
	publicKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKeyBytes,
	})
	return string(privateKeyPEM), string(publicKeyPEM), nil
}
func (q *LTIFormQueries) Get(id uint) (models.LTIForm, error) {
	var entity models.LTIForm
	result := q.First(&entity, id)
	if result.Error != nil && result.Error.Error() == "record not found" {
		return entity, utils.FiberValidationException{
			Status:    fiber.StatusNotFound,
			Exception: errors.New("LTIForm not found"),
		}
	}
	return entity, result.Error
}

func (q *LTIFormQueries) Upsert(entity *models.LTIForm) error {
	if entity == nil {
		return nil
	}
	entityDB := models.LTIForm{}
	q.Where("id = ?", entity.ID).Find(&entityDB)
	if entityDB.ID != 0 {
		// Update
		log.Debug().Msg(fmt.Sprintf("LTIFormQueries: entity update: %+v", entityDB))
		log.Info().Msg(fmt.Sprintf("LTIFormQueries: entity update id=%+v", entityDB.ID))
		entity.ID = entityDB.ID
		entity.CreatedAt = entityDB.CreatedAt
		entity.PublicKey = entityDB.PublicKey
		entity.PrivateKey = entityDB.PrivateKey
		entity.UpdatedAt = time.Now().UTC()
		result := q.Save(&entity)
		return result.Error
	}
	// Create
	log.Debug().Msg(fmt.Sprintf("LTIFormQueries: entity create: %+v", entity))
	log.Info().Msg(fmt.Sprintf("LTIFormQueries: entity create name=%+v", entity.Name))
	entity.ID = 0
	entity.CreatedAt = time.Now().UTC()
	entity.UpdatedAt = time.Now().UTC()
	privateKey, publicKey, err := ltiGenerateKeys()
	if err != nil {
		return err
	}
	entity.PublicKey = publicKey
	entity.PrivateKey = privateKey
	result := q.Create(entity)
	return result.Error
}

func (q *LTIFormQueries) List(limit int, offset int) ([]models.LTIFormListItem, error) {
	var entities []models.LTIFormListItem
	result := q.Limit(limit).Offset(offset).Order(`created_at desc`).Find(&entities)
	log.Debug().Msg(fmt.Sprintf("LTIFormQueries: entities %+v", entities))
	return entities, result.Error
}

func (q *LTIFormQueries) Delete(id uint) error {
	_ = q.Where("id = ?", id).Delete(&models.LTIForm{})
	log.Debug().Msg(fmt.Sprintf("LTIFormQueries: delete entity by id: %+v", id))
	return nil
}
