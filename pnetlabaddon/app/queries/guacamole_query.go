package queries

import (
	"gorm.io/gorm"
)

const (
	GuacamoleEntityTypeUser     = "USER"
	GuacamolePermissionTypeRead = "READ"
)

type GuacamoleQueries struct {
	*gorm.DB
}

func (q *GuacamoleQueries) UserReplace(
	userId int,
	userName string,
	passwordSha256 string,
) {
	entityId := userId + 1000
	q.Exec("REPLACE INTO `guacdb`.`guacamole_entity` (entity_id, name, type) VALUES (?, ?, ?)", entityId, userName, GuacamoleEntityTypeUser)
	q.Exec(
		"REPLACE INTO `guacdb`.`guacamole_user` (user_id, entity_id, password_hash, password_date) VALUES (?, ?, UNHEX(SHA2(?,256)), NOW())",
		entityId,
		entityId,
		passwordSha256,
	)
	q.Exec(
		"REPLACE INTO `guacdb`.`guacamole_user_permission` (entity_id, affected_user_id, permission) VALUES (?, ?, ?)",
		entityId,
		entityId,
		GuacamolePermissionTypeRead,
	)
}

func (q *GuacamoleQueries) TokenReplace(
	userId int,
	userName string,
	token string,
) {
	q.Exec("delete from `html5` where username = ? OR pod = ?", userName, userId)
	q.Exec(
		"replace into `html5` (username, pod, token) values (?, ?, ?)",
		userName,
		userId,
		token,
	)
}
