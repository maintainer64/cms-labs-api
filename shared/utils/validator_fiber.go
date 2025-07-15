package utils

import (
	fiber "github.com/gofiber/fiber/v2"
)

type FiberValidationException struct {
	Status    int
	Exception error
}

func (c FiberValidationException) Error() string {
	return c.Exception.Error()
}

type FiberSuccessResponse struct {
	Result interface{}
}

func (c FiberSuccessResponse) Error() string {
	return "Success Response"
}

func FiberValidatorBase(c *fiber.Ctx, dto interface{}) error {
	// Check, if received JSON data is valid.
	if err := c.BodyParser(dto); err != nil {
		// Return status 400 and error message.
		return FiberValidationException{fiber.StatusBadRequest, err}
	}

	// Create a new validator for a User model.
	validate := NewValidator()

	// Validate sign up fields.
	if err := validate.Struct(dto); err != nil {
		// Return, if some fields are not valid.
		return FiberValidationException{fiber.StatusBadRequest, err}
	}
	return nil
}
