package view

import (
	"context"
	commonErrors "public-transport-backend/internal/common/errors"
)

func AdminListPassengers(
	ctx context.Context,
	requestingUser *RequestingUser,
	dependencies *Dependencies,
) ([]PassengerResult, error) {
	isAdmin, err := dependencies.Repository.IsAdminUser(ctx, requestingUser)
	if !isAdmin || err != nil {
		return nil, commonErrors.NotAnAdminError()
	}

	passengers, err := dependencies.Repository.FindAll(ctx)
	if err != nil {
		return nil, commonErrors.ToGenericError(err)
	}

	return ToResults(passengers), nil
}
