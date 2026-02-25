package jsonrpc

import (
	validator "github.com/go-playground/validator/v10"
	"github.com/goccy/go-json"
	"github.com/google/uuid"
)

// NewValidator func for create a new validator for model fields.
func NewValidator() *validator.Validate {
	// Create a new validator for a Book model.
	validate := validator.New()

	// Custom validation for uuid.UUID fields.
	_ = validate.RegisterValidation("uuid", func(fl validator.FieldLevel) bool {
		field := fl.Field().String()
		if _, err := uuid.Parse(field); err != nil {
			return true
		}
		return false
	})

	return validate
}

type RpcValidatorError struct {
	Exception error
}

func (c RpcValidatorError) Error() string {
	return c.Exception.Error()
}

func ValidatorBase(c *Ctx, dto interface{}) error {
	if err := json.Unmarshal(c.Params, dto); err != nil {
		return RpcValidatorError{
			Exception: err,
		}
	}
	validate := NewValidator()
	if err := validate.Struct(dto); err != nil {
		return RpcValidatorError{Exception: err}
	}
	return nil
}
