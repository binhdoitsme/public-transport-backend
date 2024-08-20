package stubs

import (
	"context"
	"errors"
	identity "public-transport-backend/internal/features/identity/domain"
	passenger "public-transport-backend/internal/features/passenger/domain"
	"public-transport-backend/internal/features/passenger/view"
	"time"
)

// PassengerRepositoryStub is a stub implementation of the Repository interface.
type PassengerRepositoryStub struct {
	Passengers map[string]*passenger.Account
	Admins     map[uint64]bool
}

// NewPassengerRepository creates a new instance of PassengerRepositoryStub with initialized data.
func NewPassengerRepository() *PassengerRepositoryStub {
	admins := make(map[uint64]bool)
	admins[1] = true
	return &PassengerRepositoryStub{
		Passengers: make(map[string]*passenger.Account),
		Admins:     admins,
	}
}

// ExistsByPhoneNumberOrVneId checks if a passenger account exists by phone number or VneID.
func (r *PassengerRepositoryStub) ExistsByPhoneNumberOrVneId(ctx context.Context, phoneNumber string, vneId string) (bool, error) {
	for _, account := range r.Passengers {
		if account.PhoneNumber == phoneNumber || account.VneID == vneId {
			return true, nil
		}
	}
	return false, nil
}

// Save saves a passenger account and returns its ID.
func (r *PassengerRepositoryStub) Save(ctx context.Context, account *passenger.Account) (uint64, error) {
	if account == nil {
		return 0, errors.New("account cannot be nil")
	}
	r.Passengers[account.PhoneNumber] = account
	return account.Id, nil
}

// IsAdmin checks if a given user is an admin.
func (r *PassengerRepositoryStub) IsAdmin(ctx context.Context, userId uint64) (bool, error) {
	isAdmin, exists := r.Admins[userId]
	if !exists {
		return false, nil
	}
	return isAdmin, nil
}

// FindById finds a passenger account by its ID.
func (r *PassengerRepositoryStub) FindById(ctx context.Context, id uint64) (*passenger.Account, error) {
	for _, account := range r.Passengers {
		if account.Id == id {
			return account, nil
		}
	}
	return nil, nil
}

// FindAll returns a list of all passenger accounts.
func (r *PassengerRepositoryStub) FindAll(ctx context.Context, specs *view.PassengerListSpecs) ([]passenger.Account, error) {
	accounts := make([]passenger.Account, 0, len(r.Passengers))
	for _, account := range r.Passengers {
		accounts = append(accounts, *account)
	}
	return accounts, nil
}

// ------------------------------------
// AccountRepositoryStub is a stub implementation of the AccountRepository interface.
type AccountRepositoryStub struct {
	Accounts     map[uint64]*identity.Account
	TokenService *TokenServicesStub // Injected TokenServiceStub
}

// NewAccountRepositoryStub creates a new instance of AccountRepositoryStub.
func NewAccountRepository(tokenService *TokenServicesStub) *AccountRepositoryStub {
	return &AccountRepositoryStub{
		Accounts:     make(map[uint64]*identity.Account),
		TokenService: tokenService, // Injected dependency
	}
}

// ExistsByUsername checks if an account exists by username.
func (r *AccountRepositoryStub) ExistsByUsername(ctx context.Context, username string) (bool, error) {
	for _, account := range r.Accounts {
		if account.Username == username {
			return true, nil
		}
	}
	return false, nil
}

// Save stores a new account and returns its ID.
func (r *AccountRepositoryStub) Save(ctx context.Context, account *identity.Account) (uint64, error) {
	if account == nil {
		return 0, errors.New("account cannot be nil")
	}

	// Check if the account exists already
	existingAccount, exists := r.Accounts[account.Id]
	if exists {
		// Clear old refresh tokens using TokenService
		for _, token := range existingAccount.RefreshTokens {
			// Invalidate old refresh tokens
			delete(r.TokenService.refreshTokens, token.Token)
		}
	}

	// Store the updated account
	r.Accounts[account.Id] = account

	// Add new refresh tokens to TokenService and account
	for _, refreshToken := range account.RefreshTokens {
		if refreshToken.Token != "" {
			// Store the token in the TokenServiceStub
			r.TokenService.refreshTokens[refreshToken.Token] = account
		}
	}

	return account.Id, nil
}

// FindByPhoneNumberAndPassword searches for an account by phone number and password.
func (r *AccountRepositoryStub) FindByUsernameAndPassword(ctx context.Context, username string, password string) (*identity.Account, error) {
	// Iterate through the accounts and look for a match by username
	for _, account := range r.Accounts {
		if account.Username == username && account.Password == password {
			return account, nil
		}
	}

	return nil, errors.New("account not found or incorrect password")
}

// FindById retrieves an account by its ID.
func (r *AccountRepositoryStub) FindById(ctx context.Context, id uint64) (*identity.Account, error) {
	account, exists := r.Accounts[id]
	if !exists {
		return nil, errors.New("account not found")
	}
	return account, nil
}

// FindByRefreshToken uses the TokenServiceStub to find an account by refresh token.
func (r *AccountRepositoryStub) FindByRefreshToken(ctx context.Context, refreshToken string, now time.Time) (*identity.Account, error) {
	// Use the TokenService to parse the refresh token
	account, err := r.TokenService.Parse(refreshToken)
	if err != nil {
		return nil, errors.New("invalid refresh token")
	}

	// Verify the account exists in the repository
	accountInstance, exists := r.Accounts[account.Id]
	if !exists {
		return nil, errors.New("account not found")
	}
	exists = false
	for _, token := range accountInstance.RefreshTokens {
		if refreshToken == token.Token {
			exists = true
			break
		}
	}
	if !exists {
		return nil, errors.New("account not found")
	}

	return account, nil
}
