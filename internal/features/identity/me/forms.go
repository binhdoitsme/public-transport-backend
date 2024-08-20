package me

import (
	"public-transport-backend/internal/common/errors"

	"github.com/go-playground/validator"
)

type GetMyProfileForm struct {
	UserId uint64 `validate:"required"`
}

func (form *GetMyProfileForm) Validate(validate *validator.Validate) error {
	err := validate.Struct(form)
	if err != nil {
		return errors.ToValidationError(err)
	}

	return nil
}
