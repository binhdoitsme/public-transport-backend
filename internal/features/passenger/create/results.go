package create

type CreatePassengerResult struct {
	Id uint64 `json:"id" validate:"required"`
}
