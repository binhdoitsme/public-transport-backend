package view

import (
	"context"
	commonErrors "public-transport-backend/internal/common/errors"
)

func AdminListPassengers(
	ctx context.Context,
	form *PassengerListForm,
	dependencies *Dependencies,
) ([]PassengerResult, error) {
	isAdmin, err := dependencies.AdminRepository.IsAdmin(ctx, form.UserId)
	if !isAdmin || err != nil {
		return nil, commonErrors.NotAnAdminError()
	}

	passengers, err := dependencies.Repository.FindAll(ctx, &PassengerListSpecs{
		Limit:  form.PageSize,
		Offset: (form.Page - 1) * form.PageSize,
	})
	if err != nil {
		return nil, commonErrors.ToGenericError(err)
	}

	return ToResults(passengers), nil
}
