package me

import (
	"context"
	identity "public-transport-backend/internal/features/identity/domain"

	"github.com/go-playground/validator"
)

type AccountRepository interface {
	FindById(ctx context.Context, id uint64) (*identity.Account, error)
}

type TokenServices interface {
	Parse(accessToken string) (*identity.Account, error)
}

type Dependencies struct {
	AccountRepository AccountRepository
	Validate          *validator.Validate
}
