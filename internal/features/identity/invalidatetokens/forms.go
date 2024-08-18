package invalidatetokens

import (
	"github.com/go-playground/validator"
	commonErrors "public-transport-backend/internal/common/errors"
)

type InvalidateTokenForm struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
}

func (form *InvalidateTokenForm) Validate(validate *validator.Validate) error {
	err := validate.Struct(form)
	if err != nil {
		return commonErrors.ToValidationError(err)
	}
	return nil
}
