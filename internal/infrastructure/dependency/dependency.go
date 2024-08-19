package dependency

import (
	"public-transport-backend/internal/features/identity"
	"public-transport-backend/internal/features/identity/createtokens"
	"public-transport-backend/internal/features/identity/invalidatetokens"
	"public-transport-backend/internal/features/identity/me"
	"public-transport-backend/internal/features/identity/refreshtokens"
	"public-transport-backend/internal/features/identity/signup"
	"public-transport-backend/internal/features/passenger"
	"public-transport-backend/internal/features/passenger/create"
	"public-transport-backend/internal/features/passenger/view"
	"public-transport-backend/internal/infrastructure/database"
	"public-transport-backend/internal/infrastructure/database/repositories"
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

	createTokenPairDependencies     *createtokens.Dependencies
	refreshTokenPairDependencies    *refreshtokens.Dependencies
	invalidateTokenPairDependencies *invalidatetokens.Dependencies
	signupDependencies              *signup.Dependencies
	getMyProfileDependencies        *me.Dependencies
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

func (d *dependencies) InvalidateTokenPairDependenciesFactory() *invalidatetokens.Dependencies {
	return d.invalidateTokenPairDependencies
}

func (d *dependencies) SignUpDependenciesFactory() *signup.Dependencies {
	return d.signupDependencies
}

func (d *dependencies) GetMyProfileDependenciesFactory() *me.Dependencies {
	return d.getMyProfileDependencies
}

func New() Dependencies {
	validate := validator.New()
	passengerRepository := stubs.NewPassengerRepository()
	tokenService := stubs.NewTokenServices()
	db := database.New()
	accountRepository := repositories.NewAccountRepository(db.GetDB())
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
		invalidateTokenPairDependencies: &invalidatetokens.Dependencies{
			Validate:          validate,
			AccountRepository: accountRepository,
		},
		signupDependencies: &signup.Dependencies{
			Validate:         validate,
			Repository:       accountRepository,
			PasswordServices: passwordService,
		},
		getMyProfileDependencies: &me.Dependencies{
			Validate:          validate,
			AccountRepository: accountRepository,
			Tokens:            tokenService,
		},
	}
}
