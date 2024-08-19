package signup

import (
	"context"
	identity "public-transport-backend/internal/features/identity/domain"
	"public-transport-backend/internal/features/identity/support"

	"github.com/go-playground/validator"
)

type AccountRepository interface {
	ExistsByUsername(ctx context.Context, username string) (bool, error)
	Save(ctx context.Context, account *identity.Account) (uint64, error)
}

type Dependencies struct {
	Validate         *validator.Validate
	Repository       AccountRepository
	PasswordServices support.PasswordServices
}
