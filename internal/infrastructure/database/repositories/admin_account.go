package repositories

import (
	"context"
	identity "public-transport-backend/internal/features/identity/domain"
	"public-transport-backend/internal/infrastructure/database/models"
)

func (r *AccountRepositoryOnGorm) IsAdmin(ctx context.Context, userId uint64) (bool, error) {
	result := r.db.Find(
		&models.Account{ID: userId},
		"role IN ?", []string{string(identity.Admin), string(identity.SuperAdmin)})
	if err := result.Error; err != nil {
		return false, err
	}
	return result.RowsAffected > 0, result.Error
}
