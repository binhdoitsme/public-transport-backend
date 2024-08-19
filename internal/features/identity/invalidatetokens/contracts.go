package invalidatetokens

import (
	"context"
	identity "public-transport-backend/internal/features/identity/domain"
	"time"

	"github.com/go-playground/validator"
)

type AccountRepository interface {
	FindByRefreshToken(ctx context.Context, refreshToken string, now time.Time) (*identity.Account, error)
	Save(ctx context.Context, account *identity.Account) (uint64, error)
}

type Dependencies struct {
	AccountRepository AccountRepository
	Validate          *validator.Validate
}
