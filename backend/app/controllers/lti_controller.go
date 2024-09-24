package controllers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

// GetLTIInfo функция возвращает базовую информацию для интеграции LTI.
// @Description Базовая информация для интеграции LTI.
// @Summary получить базовую информацию для интеграции
// @Tags LTI
// @Accept json
// @Produce json
// @Success 200 {array} usecases.LtiFormEditOutputDTO
// @Router /v2/lti [get]
func GetLTIInfo(c *fiber.Ctx) error {
	return errors.New("error")
}

// GetLTIPublicKeys функция возвращает RSA ключи для интеграции LTI.
// @Description RSA ключи для интеграции LTI.
// @Summary получить RSA ключи для интеграции LTI
// @Tags LTI
// @Accept json
// @Produce json
// @Success 200 {array} usecases.LtiFormEditOutputDTO
// @Router /v2/lti/launch/public_keys/ [get]
func GetLTIPublicKeys(c *fiber.Ctx) error {
	return errors.New("error")
}

// LaunchLTI функция инициализирует подключение к LTI.
// @Description инициализация подключения к LTI.
// @Summary инициализация подключения к LTI
// @Tags LTI
// @Accept json
// @Produce json
// @Success 200 {array} usecases.LtiFormEditOutputDTO
// @Router /v2/lti/launch/ [post]
func LaunchLTI(c *fiber.Ctx) error {
	return errors.New("error")
}

// AuthLTI функция аутентификация пользователя по LTI.
// @Description аутентификация LTI.
// @Summary аутентификация через LTI
// @Tags LTI
// @Accept json
// @Produce json
// @Success 200 {array} usecases.LtiFormEditOutputDTO
// @Router /v2/lti/launch/authenticate/ [post]
func AuthLTI(c *fiber.Ctx) error {
	return errors.New("error")
}
