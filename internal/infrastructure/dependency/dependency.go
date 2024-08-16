package dependency

import (
	"public-transport-backend/internal/features/identity"
	"public-transport-backend/internal/features/identity/createtokens"
	"public-transport-backend/internal/features/identity/refreshtokens"
	"public-transport-backend/internal/features/identity/signup"
	"public-transport-backend/internal/features/passenger"
	"public-transport-backend/internal/features/passenger/create"
	"public-transport-backend/internal/features/passenger/view"
	"public-transport-backend/internal/infrastructure/stubs"

	"github.com/go-playground/validator"
)

type Dependencies interface {
	passenger.Dependencies
	identity.Dependencies
}

type dependencies struct {
	validate                    *validator.Validate
	createPassengerDependencies *create.Dependencies
	viewPassengerDependencies   *view.Dependencies

	createTokenPairDependencies            *createtokens.Dependencies
	refreshTokenPairDependencies *refreshtokens.Dependencies
	signupDependencies           *signup.Dependencies
}

func (d *dependencies) CreateDependenciesFactory() *create.Dependencies {
	return d.createPassengerDependencies
}

func (d *dependencies) ViewDependenciesFactory() *view.Dependencies {
	return d.viewPassengerDependencies
}

func (d *dependencies) CreateTokenPairDependenciesFactory() *createtokens.Dependencies {
	return d.createTokenPairDependencies
}

func (d *dependencies) RefreshTokenPairDependenciesFactory() *refreshtokens.Dependencies {
	return d.refreshTokenPairDependencies
}

func (d *dependencies) SignUpDependenciesFactory() *signup.Dependencies {
	return d.signupDependencies
}

func New() Dependencies {
	validate := validator.New()
	passengerRepository := stubs.NewPassengerRepository()
	tokenService := stubs.NewTokenServices()
	accountRepository := stubs.NewAccountRepository(tokenService)
	passwordService := stubs.NewPasswordServices()
	return &dependencies{
		validate: validate,
		createPassengerDependencies: &create.Dependencies{
			Validate:       validate,
			Repository:     passengerRepository,
			EventPublisher: nil,
		},
		viewPassengerDependencies: &view.Dependencies{
			Validate:   validate,
			Repository: passengerRepository,
		},

		createTokenPairDependencies: &createtokens.Dependencies{
			Validate:          validate,
			AccountRepository: accountRepository,
			Tokens:            tokenService,
			Passwords:         passwordService,
		},
		refreshTokenPairDependencies: &refreshtokens.Dependencies{
			Validate:          validate,
			AccountRepository: accountRepository,
			Tokens:            tokenService,
		},
		signupDependencies: &signup.Dependencies{
			Validate:         validate,
			Repository:       accountRepository,
			PasswordServices: passwordService,
		},
	}
}
