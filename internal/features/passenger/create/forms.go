package create

import (
	commonErrors "public-transport-backend/internal/common/errors"
	passenger "public-transport-backend/internal/features/passenger/domain"
	"time"

	"github.com/go-playground/validator"
)

type SelfPassengerForm struct {
	PhoneNumber          string                `json:"phoneNumber" validate:"required,e164"`
	VneID                string                `json:"vneId" validate:"required"`
	Name                 string                `json:"name" validate:"required"`
	DOB                  time.Time             `json:"dob" validate:"required"`
	Gender               string                `json:"gender" validate:"required"`
	PersonalImage        string                `json:"personalImage" validate:"required,url"`
	AccountType          passenger.AccountType `json:"accountType" validate:"required,oneof=Individual Student Group Elder"`
	ConfirmationDocument *string               `json:"confirmationDocument,omitempty" validate:"omitempty,url"`
}

func (form *SelfPassengerForm) Validate(validate *validator.Validate) error {
	err := validate.Struct(form)
	if err != nil {
		return commonErrors.ToValidationError(err)
	}

	return nil
}

func (form *SelfPassengerForm) ToAccount() (*passenger.Account, error) {
	return passenger.NewAccount(
		form.PhoneNumber,
		form.VneID,
		form.Name,
		form.DOB,
		form.Gender,
		form.PersonalImage,
		form.AccountType,
		form.ConfirmationDocument,
		nil,
		nil,
	)
}

type MaybeAdmin struct {
	UserId uint64
}

type AdminPassengerForm struct {
	*SelfPassengerForm
	*MaybeAdmin
}
