package dependency

import (
	"public-transport-backend/internal/features/passenger"
	"public-transport-backend/internal/features/passenger/create"
	"public-transport-backend/internal/features/passenger/view"
	"public-transport-backend/internal/infrastructure/stubs"

	"github.com/go-playground/validator"
)

type Dependencies interface {
	passenger.Dependencies
}

type dependencies struct {
	validate                    *validator.Validate
	createPassengerDependencies *create.Dependencies
	viewPassengerDependencies   *view.Dependencies
}

func (d *dependencies) CreateDependenciesFactory() *create.Dependencies {
	return d.createPassengerDependencies
}

func (d *dependencies) ViewDependenciesFactory() *view.Dependencies {
	return d.viewPassengerDependencies
}

func New() Dependencies {
	validate := validator.New()
	repository := stubs.NewPassengerRepository()
	return &dependencies{
		validate: validate,
		createPassengerDependencies: &create.Dependencies{
			Validate:       validate,
			Repository:     repository,
			EventPublisher: nil,
		},
		viewPassengerDependencies: &view.Dependencies{
			Validate:   validate,
			Repository: repository,
		},
	}
}
