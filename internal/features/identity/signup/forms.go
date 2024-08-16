package signup

import (
	"fmt"
	commonErrors "public-transport-backend/internal/common/errors"
	identity "public-transport-backend/internal/features/identity/domain"
	"public-transport-backend/internal/features/identity/support"

	"github.com/go-playground/validator"
)

type SignUpForm struct {
	Username string        `json:"username" validate:"required"`
	Password string        `json:"password" validate:"required"`
	Name     string        `json:"name" validate:"required"`
	Role     identity.Role `json:"role" validate:"oneof=User Admin"`
}

func (form *SignUpForm) Validate(validate *validator.Validate) error {
	err := validate.Struct(form)
	if err != nil {
		return commonErrors.ToValidationError(err)
	}
	if len(form.Password) < identity.MinPasswordLength {
		err = fmt.Errorf("password must contain at least %d characters", identity.MinPasswordLength)
		return commonErrors.ToValidationError(err)
	}

	return nil
}

func (form *SignUpForm) ToNewAccount(passwordService support.PasswordServices) (*identity.Account, error) {
	return identity.New(
		form.Username,
		form.Password,
		form.Name,
		form.Role,
		nil,
		nil,
	)
}
