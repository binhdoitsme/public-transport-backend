package view

import (
	"context"
	passenger "public-transport-backend/internal/features/passenger/domain"

	"github.com/go-playground/validator"
)

type PassengerListSpecs struct {
	Limit  int
	Offset int
}

type ViewPassengerRepository interface {
	FindById(ctx context.Context, id uint64) (*passenger.Account, error)
	FindAll(ctx context.Context, specs *PassengerListSpecs) ([]passenger.Account, error)
}

type AdminRepository interface {
	IsAdmin(ctx context.Context, userId uint64) (bool, error)
}

type Dependencies struct {
	Validate        *validator.Validate
	AdminRepository AdminRepository
	Repository      ViewPassengerRepository
}
