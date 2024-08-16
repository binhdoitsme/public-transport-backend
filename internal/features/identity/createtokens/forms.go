package createtokens

import (
	commonErrors "public-transport-backend/internal/common/errors"

	"github.com/go-playground/validator"
)

type NewTokensForm struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (form *NewTokensForm) Validate(validate *validator.Validate) error {
	err := validate.Struct(form)
	if err != nil {
		return commonErrors.ToValidationError(err)
	}
	return nil
}
