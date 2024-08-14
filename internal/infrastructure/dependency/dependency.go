package dependency

import (
	"public-transport-backend/internal/features/passenger"
	"public-transport-backend/internal/features/passenger/create"
	"public-transport-backend/internal/infrastructure/stubs"

	"github.com/go-playground/validator"
)

type Dependencies interface {
	passenger.Dependencies
}

type dependencies struct {
	validate                    *validator.Validate
	createPassengerDependencies *create.Dependencies
}

func (d *dependencies) CreateDependenciesFactory() *create.Dependencies {
	return d.createPassengerDependencies
}

func New() Dependencies {
	validate := validator.New()
	return &dependencies{
		validate: validate,
		createPassengerDependencies: &create.Dependencies{
			Validate:       validate,
			Repository:     stubs.NewPassengerRepository(),
			EventPublisher: nil,
		},
	}
}
