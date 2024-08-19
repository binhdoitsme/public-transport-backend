package createtokens

import (
	"context"
	identity "public-transport-backend/internal/features/identity/domain"

	"github.com/go-playground/validator"
)

type AccountRepository interface {
	FindByUsernameAndPassword(ctx context.Context, phoneNumber string, password string) (*identity.Account, error)
	Save(ctx context.Context, account *identity.Account) (uint64, error)
}

type TokenServices interface {
	NewRefreshToken(ctx context.Context, account *identity.Account) (string, error)
	NewAccessToken(ctx context.Context, account *identity.Account, refreshToken string) (string, error)
	Parse(accessToken string) (*identity.Account, error)
}

type PasswordServices interface {
	ToStoredForm(ctx context.Context, password string) (string, error)
}

type Dependencies struct {
	AccountRepository AccountRepository
	Tokens            TokenServices
	Passwords         PasswordServices
	Validate          *validator.Validate
}
