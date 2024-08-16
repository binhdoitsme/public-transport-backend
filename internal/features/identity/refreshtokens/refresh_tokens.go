package refreshtokens

import (
	"context"

	commonErrors "public-transport-backend/internal/common/errors"
)

func RefreshTokenPair(
	ctx context.Context,
	form *RefreshTokenForm,
	dependencies *Dependencies,
) (*SessionResult, error) {
	if err := form.Validate(dependencies.Validate); err != nil {
		return nil, err
	}

	accountRepository := dependencies.AccountRepository

	account, err := accountRepository.FindByRefreshToken(ctx, form.RefreshToken)

	if err != nil {
		return nil, commonErrors.ToGenericError(err)
	}

	if account == nil {
		return &SessionResult{}, nil
	}
	accessToken, err := dependencies.Tokens.NewAccessToken(ctx, account, form.RefreshToken)
	if err != nil {
		return nil, commonErrors.ToGenericError(err)
	}

	return &SessionResult{
		Ok: true, RefreshToken: form.RefreshToken, AccessToken: accessToken,
	}, nil
}
