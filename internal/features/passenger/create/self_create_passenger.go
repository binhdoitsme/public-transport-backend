// contains Forms-Handler(service-repository-event publisher)-Result
package create

import (
	"context"
	commonErrors "public-transport-backend/internal/common/errors"
	passenger "public-transport-backend/internal/features/passenger/domain"
)

func SelfCreatePassenger(
	ctx context.Context,
	form *SelfPassengerForm,
	dependencies *Dependencies,
) (*CreatePassengerResult, error) {
	err := form.Validate(dependencies.Validate)
	if err != nil {
		return nil, err
	}

	existed, err := dependencies.Repository.ExistsByPhoneNumberOrVneId(ctx, form.PhoneNumber, form.VneID)

	if err != nil {
		return nil, commonErrors.ToGenericError(err)
	}

	if existed {
		return nil, alreadyExistsError(form.PhoneNumber, form.VneID)
	}

	account, err := form.ToAccount()

	if err != nil {
		return nil, err
	}

	if account.AccountType == passenger.Individual {
		// individual pass does not need to be approved
		account.Status = passenger.Approved
	}

	id, err := dependencies.Repository.Save(ctx, account)

	if err != nil {
		return nil, commonErrors.ToGenericError(err)
	}

	if account.AccountType != passenger.Individual {
		err = dependencies.EventPublisher.RequestApproval(id)
		if err != nil {
			return nil, commonErrors.ToGenericError(err)
		}
	}

	return &CreatePassengerResult{Id: id}, nil
}
