package create

import (
	"context"
	"github.com/go-playground/validator"
	passenger "public-transport-backend/internal/features/passenger/domain"
)

type CreatePassengerEventPublisher interface {
	RequestApproval(id uint64) error
}

type CreatePassengerRepository interface {
	ExistsByPhoneNumberOrVneId(ctx context.Context, phoneNumber string, vneId string) (bool, error)
	Save(ctx context.Context, account *passenger.Account) (uint64, error)
}

type AdminRepository interface {
	IsAdmin(ctx context.Context, maybeAdmin *MaybeAdmin) (bool, error)
}

type Repository interface {
	CreatePassengerRepository
	AdminRepository
}

type Dependencies struct {
	Validate       *validator.Validate
	Repository     Repository
	EventPublisher CreatePassengerEventPublisher
}
