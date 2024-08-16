package createtokens

import (
	"context"
	commonErrors "public-transport-backend/internal/common/errors"
)

func NewTokenPair(ctx context.Context, form *NewTokensForm, dependencies *Dependencies) (*SessionResult, error) {
	if err := form.Validate(dependencies.Validate); err != nil {
		return nil, err
	}
	maybeStoredPassword, err := dependencies.Passwords.ToStoredForm(ctx, form.Password)

	if err != nil {
		return nil, commonErrors.ToGenericError(err)
	}

	account, err := dependencies.AccountRepository.FindByUsernameAndPassword(ctx, form.Username, maybeStoredPassword)

	if err != nil {
		return nil, commonErrors.ToGenericError(err)
	}

	if account == nil {
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

	return &SessionResult{
		Ok: true, RefreshToken: refreshToken, AccessToken: accessToken,
	}, nil
}
