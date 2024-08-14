package create

import (
	"context"
	commonErrors "public-transport-backend/internal/common/errors"
	passenger "public-transport-backend/internal/features/passenger/domain"
)

func AdminCreatePassenger(
	ctx context.Context,
	form *AdminPassengerForm,
	dependencies *Dependencies,
) (*CreatePassengerResult, error) {
	err := form.Validate(dependencies.Validate)
	if err != nil {
		return nil, err
	}

	isAdmin, err := dependencies.Repository.IsAdmin(ctx, form.MaybeAdmin)
	if err != nil {
		return nil, commonErrors.ToGenericError(err)
	}

	if !isAdmin {
		return nil, commonErrors.NotAnAdminError(form.UserId)
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

	account.Status = passenger.Approved
	id, err := dependencies.Repository.Save(ctx, account)

	if err != nil {
		return nil, commonErrors.ToGenericError(err)
	}

	return &CreatePassengerResult{Id: id}, nil
}
