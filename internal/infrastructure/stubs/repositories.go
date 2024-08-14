package stubs

import (
	"context"
	"errors"
	"public-transport-backend/internal/features/passenger/create"
	passenger "public-transport-backend/internal/features/passenger/domain"
)

// PassengerRepositoryStub is a stub implementation of the Repository interface.
type PassengerRepositoryStub struct {
	// You can add fields here to store in-memory data if needed
	Passengers map[string]*passenger.Account
	Admins     map[uint64]bool
}

// NewPassengerRepository creates a new instance of RepositoryStub with initialized data.
func NewPassengerRepository() *PassengerRepositoryStub {
	return &PassengerRepositoryStub{
		Passengers: make(map[string]*passenger.Account),
		Admins:     make(map[uint64]bool),
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
