package repositories

import (
	"context"
	passenger "public-transport-backend/internal/features/passenger/domain"
	"public-transport-backend/internal/features/passenger/view"
	"public-transport-backend/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type PassengerRepositoryOnGorm struct {
	db *gorm.DB
}

func NewPassengerRepository(db *gorm.DB) *PassengerRepositoryOnGorm {
	return &PassengerRepositoryOnGorm{db}
}

// ExistsByPhoneNumberOrVneId checks if a passenger account exists by phone number or VneID.
func (r *PassengerRepositoryOnGorm) ExistsByPhoneNumberOrVneId(
	ctx context.Context,
	phoneNumber string,
	vneId string,
) (bool, error) {
	result := r.db.WithContext(ctx).Limit(1).Find(&models.Passenger{
		PhoneNumber: phoneNumber,
		VneID:       vneId,
	})
	if err := result.Error; err != nil {
		return false, err
	}
	return result.RowsAffected > 0, nil
}

// Save saves a passenger account and returns its ID.
func (r *PassengerRepositoryOnGorm) Save(
	ctx context.Context,
	account *passenger.Account,
) (uint64, error) {
	record := &models.Passenger{
		ID:                   account.Id,
		PhoneNumber:          account.PhoneNumber,
		VneID:                account.VneID,
		Name:                 account.Name,
		DOB:                  account.DOB,
		Gender:               account.Gender,
		PersonalImage:        account.PersonalImage,
		ConfirmationDocument: account.ConfirmationDocument,
		AccountType:          string(account.AccountType),
		Status:               string(account.Status),
	}
	result := r.db.WithContext(ctx).Where(models.Passenger{ID: account.Id}).Assign(record).FirstOrCreate(record)
	if err := result.Error; err != nil {
		return 0, err
	}
	return account.Id, nil
}

// FindById finds a passenger account by its ID.
func (r *PassengerRepositoryOnGorm) FindById(ctx context.Context, id uint64) (*passenger.Account, error) {
	passenger := &models.Passenger{ID: id}
	result := r.db.WithContext(ctx).First(passenger)
	if err := result.Error; err != nil {
		return nil, err
	}
	return passenger.ToPassenger(), nil
}

// FindAll returns a list of all passenger accounts.
func (r *PassengerRepositoryOnGorm) FindAll(ctx context.Context, specs *view.PassengerListSpecs) ([]passenger.Account, error) {
	passengers := make([]models.Passenger, 0)
	result := r.db.WithContext(ctx).Limit(specs.Limit).Offset(specs.Offset).Find(passengers)
	if err := result.Error; err != nil {
		return nil, err
	}
	passengerAccounts := make([]passenger.Account, 0)
	for _, p := range passengers {
		passengerAccounts = append(passengerAccounts, *p.ToPassenger())
	}
	return passengerAccounts, nil
}
