package repositories

import (
	"context"
	identity "public-transport-backend/internal/features/identity/domain"
	"public-transport-backend/internal/infrastructure/database/models"
	"time"

	"gorm.io/gorm"
)

// ------------------------------------
type AccountRepositoryOnGorm struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *AccountRepositoryOnGorm {
	return &AccountRepositoryOnGorm{db}
}

// Save stores a new account and returns its ID.
func (r *AccountRepositoryOnGorm) Save(ctx context.Context, account *identity.Account) (uint64, error) {
	refreshTokens := make([]models.RefreshToken, 0, len(account.RefreshTokens))
	refreshTokenStrings := make([]string, 0, len(account.RefreshTokens))
	for _, token := range account.RefreshTokens {
		refreshTokens = append(
			refreshTokens,
			models.RefreshToken{
				AccountID:  account.Id,
				Token:      token.Token,
				Expiration: token.Expiration,
			})
		refreshTokenStrings = append(refreshTokenStrings, token.Token)
	}
	accountRecord := models.Account{
		ID:            account.Id,
		Username:      account.Username,
		Password:      account.Password,
		Name:          account.Name,
		Role:          string(account.Role),
		PersonalImage: account.PersonalImage,
		RefreshTokens: refreshTokens,
	}
	result := r.db.WithContext(ctx).Delete(&models.RefreshToken{}, "token NOT IN ?", refreshTokenStrings)
	if err := result.Error; err != nil {
		return 0, err
	}
	result = r.db.WithContext(ctx).Save(refreshTokens)
	if err := result.Error; err != nil {
		return 0, err
	}
	result = r.db.WithContext(ctx).Where(models.Account{ID: account.Id}).Assign(accountRecord).FirstOrCreate(&accountRecord)
	if err := result.Error; err != nil {
		return 0, err
	}
	return account.Id, nil
}

// ExistsByUsername checks if an account exists by username.
func (r *AccountRepositoryOnGorm) ExistsByUsername(ctx context.Context, username string) (bool, error) {
	account := &models.Account{Username: username}
	result := r.db.WithContext(ctx).Limit(1).Find(account)
	if err := result.Error; err != nil {
		return false, err
	}

	return result.RowsAffected > 0, nil
}

// FindByPhoneNumberAndPassword searches for an account by phone number and password.
func (r *AccountRepositoryOnGorm) FindByUsername(ctx context.Context, username string) (*identity.Account, error) {
	account := &models.Account{Username: username}
	result := r.db.WithContext(ctx).Limit(1).Preload("RefreshTokens").Find(account)
	if err := result.Error; err != nil {
		return nil, err
	}
	if result.RowsAffected <= 0 {
		return nil, nil
	}

	return account.ToAccount(), nil
}

// FindById retrieves an account by its ID.
func (r *AccountRepositoryOnGorm) FindById(ctx context.Context, id uint64) (*identity.Account, error) {
	account := &models.Account{ID: id}
	result := r.db.WithContext(ctx).Limit(1).Preload("RefreshTokens").Find(account)
	if err := result.Error; err != nil {
		return nil, err
	}
	if result.RowsAffected <= 0 {
		return nil, nil
	}

	return account.ToAccount(), nil
}

// FindByRefreshToken uses the TokenServiceStub to find an account by refresh token.
func (r *AccountRepositoryOnGorm) FindByRefreshToken(ctx context.Context, refreshToken string, now time.Time) (*identity.Account, error) {
	token := &models.RefreshToken{Token: refreshToken}
	result := r.db.WithContext(ctx).Limit(1).Find(token, "expiration > ?", now)
	if err := result.Error; err != nil {
		return nil, err
	}
	if result.RowsAffected <= 0 {
		return nil, nil
	}
	account := &models.Account{ID: token.AccountID}
	result = r.db.WithContext(ctx).Limit(1).Preload("RefreshTokens").Find(account)
	if err := result.Error; err != nil {
		return nil, err
	}
	return account.ToAccount(), nil
}
