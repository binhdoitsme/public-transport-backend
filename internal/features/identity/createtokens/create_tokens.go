package createtokens

import (
	"context"
	commonErrors "public-transport-backend/internal/common/errors"
)

func NewTokenPair(ctx context.Context, form *NewTokensForm, dependencies *Dependencies) (*SessionResult, error) {
	if err := form.Validate(dependencies.Validate); err != nil {
		return nil, err
	}
	account, err := dependencies.AccountRepository.FindByUsername(ctx, form.Username)

	if err != nil {
		return nil, commonErrors.ToGenericError(err)
	}

	if account == nil {
		return &SessionResult{}, nil
	}

	isCorrectPassword := dependencies.Passwords.Compare(ctx, account.Password, form.Password)
	if !isCorrectPassword {
		return &SessionResult{}, nil
	}

	// create new access/refresh token pair
	refreshToken, err := dependencies.Tokens.NewRefreshToken(ctx, account)
	if err != nil {
		return nil, commonErrors.ToGenericError(err)
	}
	accessToken, err := dependencies.Tokens.NewAccessToken(ctx, account, refreshToken)
	if err != nil {
		return nil, commonErrors.ToGenericError(err)
	}
	account.AddRefreshToken(refreshToken)
	_, err = dependencies.AccountRepository.Save(ctx, account)
	if err != nil {
		return nil, commonErrors.ToGenericError(err)
	}

	return &SessionResult{
		Ok: true, RefreshToken: refreshToken, AccessToken: accessToken,
	}, nil
}
