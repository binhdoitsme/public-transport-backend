package view

import (
	"context"
	commonErrors "public-transport-backend/internal/common/errors"
)

func AdminViewPassenger(
	ctx context.Context,
	form *AdminPassengerByIdForm,
	dependencies *Dependencies,
) (*PassengerResult, error) {
	isAdmin, err := dependencies.Repository.IsAdminUser(ctx, form.RequestingUser)
	if !isAdmin || err != nil {
		return nil, commonErrors.NotAnAdminError()
	}

	passenger, err := dependencies.Repository.FindById(ctx, form.Id)
	if err != nil {
		return nil, commonErrors.ToGenericError(err)
	}

	if passenger == nil {
		return nil, notFound(form.Id)
	}

	return ToResult(passenger), err
}

func ViewMyPassenger(
	ctx context.Context,
	requestingUser *RequestingUser,
	dependencies *Dependencies,
) (*PassengerResult, error) {
	passenger, err := dependencies.Repository.FindById(ctx, requestingUser.UserId)
	if err != nil {
		return nil, commonErrors.ToGenericError(err)
	}
	return ToResult(passenger), err
}
