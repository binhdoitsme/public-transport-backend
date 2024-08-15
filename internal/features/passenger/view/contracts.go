package view

import (
	"context"
	passenger "public-transport-backend/internal/features/passenger/domain"

	"github.com/go-playground/validator"
)

type ViewPassengerRepository interface {
	FindById(ctx context.Context, id uint64) (*passenger.Account, error)
	FindByUserId(ctx context.Context, userId uint64) (*passenger.Account, error)
	FindAll(ctx context.Context) ([]passenger.Account, error)
}

type AdminRepository interface {
	IsAdminUser(ctx context.Context, requestingUser *RequestingUser) (bool, error)
}

type Repository interface {
	ViewPassengerRepository
	AdminRepository
}

type Dependencies struct {
	Validate   *validator.Validate
	Repository Repository
}
