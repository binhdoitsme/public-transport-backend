package view

import (
	commonErrors "public-transport-backend/internal/common/errors"

	"github.com/go-playground/validator"
)

type RequestingUser struct {
	UserId uint64
}

type PassengerByIdForm struct {
	Id uint64
}

type AdminPassengerByIdForm struct {
	Id uint64
	*RequestingUser
}

type PassengerListForm struct {
	*RequestingUser
	Page     int `validator:"gte=1"`
	PageSize int `validator:"gte=1,lte=50"`
}

func (form *PassengerListForm) Validate(validate *validator.Validate) error {
	err := validate.Struct(form)
	if err != nil {
		return commonErrors.ToValidationError(err)
	}
	return nil
}
