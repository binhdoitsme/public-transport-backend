package me

import (
	"context"
	commonErrors "public-transport-backend/internal/common/errors"
)

func GetMyProfile(
	ctx context.Context,
	form *GetMyProfileForm,
	dependencies *Dependencies,
) (*ProfileResult, error) {
	if err := form.Validate(dependencies.Validate); err != nil {
		return nil, err
	}

	account, err := dependencies.AccountRepository.FindById(ctx, form.UserId)

	if err != nil {
		return nil, commonErrors.ToGenericError(err)
	}

	if account == nil {
		return nil, nil
	}

	return &ProfileResult{
		Name:          account.Name,
		PersonalImage: account.PersonalImage,
	}, nil
}
