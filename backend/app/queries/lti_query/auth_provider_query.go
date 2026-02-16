package lti_query

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"time"

	"github.com/rs/zerolog"

	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/shared/jsonrpc"
	"gorm.io/gorm"
)

type AuthProviderQueries struct {
	DB     *gorm.DB
	Logger *zerolog.Logger
}

func ltiGenerateKeys(l *zerolog.Logger) (string, string, error) {
	l.Info().Msg("LTIGenerate RSA keys")
	// Generate RSA private key
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		l.Warn().Msg(fmt.Sprintf("Error generating private key: %+v", err))
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
		l.Warn().Msg(fmt.Sprintf("Error marshaling public key: %+v", err))
		return "", "", err
	}
	publicKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKeyBytes,
	})
	return string(privateKeyPEM), string(publicKeyPEM), nil
}
func (q *AuthProviderQueries) Get(id uint) (models.AuthProvider, error) {
	var entity models.AuthProvider
	result := q.DB.First(&entity, id)
	if result.Error != nil && result.Error.Error() == "record not found" {
		return entity, jsonrpc.NewRpcError("lti_form_not_found", "lti form not found")
	}
	return entity, result.Error
}

func (q *AuthProviderQueries) Upsert(entity *models.AuthProvider) error {
	if entity == nil {
		return nil
	}
	entityDB := models.AuthProvider{}
	q.DB.Where("id = ?", entity.ID).Find(&entityDB)
	if entityDB.ID != 0 {
		// Update
		q.Logger.Debug().Msg(fmt.Sprintf("AuthProviderQueries: entity update: %+v", entityDB))
		q.Logger.Info().Msg(fmt.Sprintf("AuthProviderQueries: entity update id=%+v", entityDB.ID))
		entity.ID = entityDB.ID
		entity.Type = entityDB.Type
		entity.CreatedAt = entityDB.CreatedAt
		entity.PublicKey = entityDB.PublicKey
		entity.PrivateKey = entityDB.PrivateKey
		entity.UpdatedAt = time.Now().UTC()
		result := q.DB.Save(&entity)
		return result.Error
	}
	// Create
	q.Logger.Debug().Msg(fmt.Sprintf("AuthProviderQueries: entity create: %+v", entity))
	q.Logger.Info().Msg(fmt.Sprintf("AuthProviderQueries: entity create name=%+v", entity.Name))
	entity.ID = 0
	entity.CreatedAt = time.Now().UTC()
	entity.UpdatedAt = time.Now().UTC()
	privateKey, publicKey, err := ltiGenerateKeys(q.Logger)
	if err != nil {
		return err
	}
	if entity.Type == models.AuthProviderTypeLTI {
		entity.PublicKey = publicKey
		entity.PrivateKey = privateKey
	}
	result := q.DB.Create(entity)
	return result.Error
}

const MaxLimitCount = 5000

func (q *AuthProviderQueries) List(
	search string,
	limit int,
	offset int,
) ([]models.AuthProviderListItem, int64, error) {
	var entities []models.AuthProviderListItem
	result := q.listFilter(
		search,
		q.DB.Limit(MaxLimitCount).Offset(0),
	).Find(&entities)
	count := result.RowsAffected
	q.Logger.Debug().Msg(fmt.Sprintf("AuthProviderQueries list: count %+v", count))
	result = q.listFilter(
		search,
		q.DB.Limit(limit).Offset(offset),
	).Find(&entities)
	q.Logger.Debug().Msg(fmt.Sprintf("AuthProviderQueries list: entities %+v", entities))
	return entities, count, result.Error
}

func (q *AuthProviderQueries) listFilter(search string, tx *gorm.DB) *gorm.DB {
	tx = tx.Model(&models.AuthProvider{})
	tx = tx.Order(`created_at desc`)
	if search == "" {
		return tx
	}
	tx = tx.Where("base_uri LIKE ?", fmt.Sprintf("%%%s%%", search))
	tx = tx.Or("id = ?", search)
	return tx
}

func (q *AuthProviderQueries) SSOURLList() ([]models.AuthProviderListItem, error) {
	var entities []models.AuthProviderListItem
	result := q.DB.Model(&models.AuthProvider{}).Order(
		`created_at desc`,
	).Where(
		`sso_url != '' AND sso_url is not null`,
	).Limit(MaxLimitCount).Offset(0).Find(&entities)
	count := result.RowsAffected
	q.Logger.Debug().Msg(fmt.Sprintf("SSOURLList list: count %+v", count))
	return entities, result.Error
}

func (q *AuthProviderQueries) Delete(id uint) error {
	err := q.DB.Where("id = ?", id).Delete(&models.AuthProvider{}).Error
	q.Logger.Debug().Msg(fmt.Sprintf("AuthProviderQueries: delete entity by id: %+v", id))
	return err
}
