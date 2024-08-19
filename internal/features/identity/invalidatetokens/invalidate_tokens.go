package invalidatetokens

import (
	"context"
	commonErrors "public-transport-backend/internal/common/errors"
)

func InvalidateToken(
	ctx context.Context,
	form *InvalidateTokenForm,
	dependencies *Dependencies,
) (*InvalidateTokenResult, error) {
	if err := form.Validate(dependencies.Validate); err != nil {
		return nil, err
	}

	accountRepository := dependencies.AccountRepository

	account, err := accountRepository.FindByRefreshToken(ctx, form.RefreshToken)

	if err != nil {
		return nil, commonErrors.ToGenericError(err)
	}

	if account == nil {
		return &InvalidateTokenResult{}, nil
	}
	// invalidate
	account.InvalidateToken(form.RefreshToken)
	accountRepository.Save(ctx, account)

	return &InvalidateTokenResult{Ok: true}, nil
}
