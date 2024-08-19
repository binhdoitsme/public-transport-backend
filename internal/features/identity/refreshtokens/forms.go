package refreshtokens

import (
	commonErrors "public-transport-backend/internal/common/errors"
	"time"

	"github.com/go-playground/validator"
)

type RefreshTokenForm struct {
	RefreshToken string    `json:"refreshToken" validate:"required"`
	Now          time.Time `json:"-" validate:"required"`
}

func (form *RefreshTokenForm) Validate(validate *validator.Validate) error {
	err := validate.Struct(form)
	if err != nil {
		return commonErrors.ToValidationError(err)
	}
	return nil
}
