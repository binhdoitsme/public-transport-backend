package signup

import (
	"context"
	commonErrors "public-transport-backend/internal/common/errors"
)

func CreateUserAccount(
	ctx context.Context,
	form *SignUpForm,
	dependencies *Dependencies,
) (*SignUpResult, error) {
	err := form.Validate(dependencies.Validate)
	if err != nil {
		return nil, err
	}

	repository := dependencies.Repository
	exists, err := repository.ExistsByUsername(form.Username)

	if err != nil {
		return nil, commonErrors.ToGenericError(err)
	}
	if exists {
		return nil, duplicateUsername(form.Username)
	}

	account, err := form.ToNewAccount(dependencies.PasswordServices)
	if err != nil {
		return nil, err
	}

	id, err := repository.Save(account)
	if err != nil {
		return nil, commonErrors.ToGenericError(err)
	}
	return &SignUpResult{id}, nil
}
