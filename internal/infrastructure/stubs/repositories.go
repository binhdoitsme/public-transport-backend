package stubs

import (
	"context"
	"errors"
	"public-transport-backend/internal/features/passenger/create"
	passenger "public-transport-backend/internal/features/passenger/domain"
	"public-transport-backend/internal/features/passenger/view"
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
func (r *PassengerRepositoryStub) IsAdmin(ctx context.Context, maybeAdmin *create.MaybeAdmin) (bool, error) {
	if maybeAdmin == nil {
		return false, nil
	}
	isAdmin, exists := r.Admins[maybeAdmin.UserId]
	if !exists {
		return false, nil
	}
	return isAdmin, nil
}

// IsAdmin checks if a given user is an admin.
func (r *PassengerRepositoryStub) IsAdminUser(ctx context.Context, requestingUser *view.RequestingUser) (bool, error) {
	return r.IsAdmin(ctx, &create.MaybeAdmin{UserId: requestingUser.UserId})
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

// FindByUserId finds a passenger account by its User ID.
func (r *PassengerRepositoryStub) FindByUserId(ctx context.Context, userId uint64) (*passenger.Account, error) {
	for _, account := range r.Passengers {
		if account.Id == userId {
			return account, nil
		}
	}
	return nil, errors.New("account not found")
}

// FindAll returns a list of all passenger accounts.
func (r *PassengerRepositoryStub) FindAll(ctx context.Context) ([]passenger.Account, error) {
	accounts := make([]passenger.Account, 0, len(r.Passengers))
	for _, account := range r.Passengers {
		accounts = append(accounts, *account)
	}
	return accounts, nil
}
