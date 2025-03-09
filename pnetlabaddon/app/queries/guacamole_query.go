package queries

import (
	"fmt"

	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

const (
	GuacamoleEntityTypeUser     = "USER"
	GuacamolePermissionTypeRead = "READ"
)

type GuacamoleQueries struct {
	*gorm.DB
	*zerolog.Logger
}

func (q *GuacamoleQueries) UserReplace(
	userId int,
	userName string,
	passwordSha256 string,
) error {
	q.Logger.Info().Msg(fmt.Sprintf("GuacamoleQueries: UserReplace by id: %+v", userId))
	entityId := userId + 1000
	err := q.Exec(
		"REPLACE INTO `guacdb`.`guacamole_entity` (entity_id, name, type) VALUES (?, ?, ?)",
		entityId,
		userName,
		GuacamoleEntityTypeUser,
	).Error
	if err != nil {
		return err
	}
	err = q.Exec(
		"REPLACE INTO `guacdb`.`guacamole_user` (user_id, entity_id, password_hash, password_date) VALUES (?, ?, UNHEX(SHA2(?,256)), NOW())",
		entityId,
		entityId,
		passwordSha256,
	).Error
	if err != nil {
		return err
	}
	err = q.Exec(
		"REPLACE INTO `guacdb`.`guacamole_user_permission` (entity_id, affected_user_id, permission) VALUES (?, ?, ?)",
		entityId,
		entityId,
		GuacamolePermissionTypeRead,
	).Error
	return err
}

func (q *GuacamoleQueries) TokenReplace(
	userId int,
	userName string,
	token string,
) error {
	err := q.Exec("delete from `html5` where username = ? OR pod = ?", userName, userId).Error
	if err != nil {
		return err
	}
	err = q.Exec(
		"replace into `html5` (username, pod, token) values (?, ?, ?)",
		userName,
		userId,
		token,
	).Error
	return err
}
