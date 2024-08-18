package invalidatetokens

import (
	"context"
	identity "public-transport-backend/internal/features/identity/domain"

	"github.com/go-playground/validator"
)

type AccountRepository interface {
	FindByRefreshToken(ctx context.Context, refreshToken string) (*identity.Account, error)
	Save(account *identity.Account) (uint64, error)
}

type Dependencies struct {
	AccountRepository AccountRepository
	Validate          *validator.Validate
}
