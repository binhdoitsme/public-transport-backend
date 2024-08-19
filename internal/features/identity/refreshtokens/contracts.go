package refreshtokens

import (
	"context"
	identity "public-transport-backend/internal/features/identity/domain"
	"time"

	"github.com/go-playground/validator"
)

type AccountRepository interface {
	FindByRefreshToken(ctx context.Context, refreshToken string, now time.Time) (*identity.Account, error)
}

type TokenServices interface {
	NewAccessToken(ctx context.Context, account *identity.Account, refreshToken string) (string, error)
	Parse(accessToken string) (*identity.Account, error)
}

type Dependencies struct {
	AccountRepository AccountRepository
	Tokens            TokenServices
	Validate          *validator.Validate
}
