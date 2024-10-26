package queries

import (
	"errors"
	"fmt"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/ory/go-convenience/mapx"
	"github.com/ory/go-convenience/stringsx"
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/pkg/utils"
	"gorm.io/gorm"
	"strings"
	"time"
)

var (
	UserNotFoundError = errors.New("User not found")
	UserNotActive     = errors.New("User is deactivated")
)

type UserQueries struct {
	*gorm.DB
}

func (q *UserQueries) Get(id uint) (models.User, error) {
	var entity models.User
	result := q.First(&entity, id)
	if result.Error != nil && result.Error.Error() == "record not found" {
		return entity, utils.FiberValidationException{
			Status:    fiber.StatusNotFound,
			Exception: UserNotFoundError,
		}
	}
	if !entity.IsActive() {
		return entity, utils.FiberValidationException{
			Status:    fiber.StatusForbidden,
			Exception: UserNotActive,
		}
	}
	return entity, result.Error
}

func (q *UserQueries) Upsert(entity *models.User) error {
	if entity == nil {
		return nil
	}
	entityDB := models.User{}
	q.Where("id = ?", entity.ID).Find(&entityDB)
	if entityDB.ID != 0 {
		// Update
		log.Debug().Msg(fmt.Sprintf("UserQueries: entity update: %+v", entityDB))
		log.Info().Msg(fmt.Sprintf("UserQueries: entity update id=%+v", entityDB.ID))
		entity.ID = entityDB.ID
		entity.CreatedAt = entityDB.CreatedAt
		entity.UpdatedAt = time.Now().UTC()
		result := q.Save(&entity)
		return result.Error
	} else {
		// Create
		log.Debug().Msg(fmt.Sprintf("UserQueries: entity create: %+v", entity))
		log.Info().Msg(fmt.Sprintf("UserQueries: entity create email=%+v", entity.Email))
		entity.ID = 0
		entity.CreatedAt = time.Now().UTC()
		entity.UpdatedAt = time.Now().UTC()
		result := q.Create(entity)
		return result.Error
	}
}

func (q *UserQueries) GetByEmail(email string) (models.User, error) {
	var entity models.User
	err := utils.FiberValidationException{
		Status:    fiber.StatusNotFound,
		Exception: UserNotFoundError,
	}
	log.Info().Msg(fmt.Sprintf("UserQueries: get user by email=%+v", email))
	if email == "" {
		return entity, err
	}
	q.Where("email = ?", email).Limit(1).Find(&entity)
	if entity.Email != email {
		return entity, err
	}
	if !entity.IsActive() {
		log.Info().Msg(fmt.Sprintf("UserQueries: user is not active by email=%+v", email))
		return entity, utils.FiberValidationException{
			Status:    fiber.StatusForbidden,
			Exception: UserNotActive,
		}
	}
	return entity, nil
}

func (q *UserQueries) GetByLaunchID(launchID string) (models.User, error) {
	var entity models.User
	err := utils.FiberValidationException{
		Status:    fiber.StatusNotFound,
		Exception: UserNotFoundError,
	}
	log.Info().Msg(fmt.Sprintf("UserQueries: get user by launch_id=%+v", launchID))
	if launchID == "" {
		return entity, err
	}
	q.Where("launch_id = ?", launchID).Limit(1).Find(&entity)
	if entity.LastLaunchID != launchID {
		return entity, err
	}
	if !entity.IsActive() {
		log.Info().Msg(fmt.Sprintf("UserQueries: user is not active by launchID=%+v", launchID))
		return entity, utils.FiberValidationException{
			Status:    fiber.StatusForbidden,
			Exception: UserNotActive,
		}
	}
	return entity, nil
}

func (q *UserQueries) List(
	search string,
	limit int,
	offset int,
) ([]models.UserListItem, int64, error) {
	var entities []models.UserListItem
	result := q.listFilter(
		search,
		q.Limit(MaxLimitCount).Offset(0),
	).Find(&entities)
	count := result.RowsAffected
	log.Debug().Msg(fmt.Sprintf("UserQueries: count %+v", count))
	result = q.listFilter(
		search,
		q.Limit(limit).Offset(offset),
	).Find(&entities)
	log.Debug().Msg(fmt.Sprintf("UserQueries: entities %+v", entities))
	return entities, count, result.Error
}

func (q *UserQueries) listFilter(search string, tx *gorm.DB) *gorm.DB {
	tx = tx.Order(`created_at desc`)
	if search == "" {
		return tx
	}
	tx = tx.Where("email LIKE ?", fmt.Sprintf("%%%s%%", search))
	tx = tx.Or("name LIKE ?", fmt.Sprintf("%%%s%%", search))
	return tx
}

func (q *UserQueries) Delete(id uint) error {
	_ = q.Where("id = ?", id).Update("deleted_at", time.Now().UTC())
	log.Debug().Msg(fmt.Sprintf("UserQueries: delete entity by id: %+v", id))
	return nil
}

func (q *UserQueries) UpdateByLaunchData(
	launchID string,
	launchData json.RawMessage,
) error {
	log.Info().Msg(fmt.Sprintf("UserQueries: update launch data by id: %+v", launchID))
	var jwtTokenPayload map[interface{}]interface{}
	if err := json.Unmarshal(launchData, &jwtTokenPayload); err != nil {
		return err
	}
	email := strings.ToLower(mapx.GetStringDefault(jwtTokenPayload, "email", ""))
	userFromDB, _ := q.GetByEmail(email)
	user := models.User{}
	user.ID = userFromDB.ID
	user.Email = email
	user.Name = mapx.GetStringDefault(jwtTokenPayload, "name", "")
	user.UserRole = stringsx.Coalesce(userFromDB.UserRole, models.UsersRoleStudent)
	user.GroupName = stringsx.Coalesce(user.GroupName, userFromDB.GroupName)
	user.LTIUserID = mapx.GetStringDefault(jwtTokenPayload, "sub", "")
	user.DeletedAt = userFromDB.DeletedAt
	user.LastLaunchID = launchID
	err := q.Upsert(&user)
	return err
}
