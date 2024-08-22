package passengerhub

type PassengerEventType string

const (
	PassengerCreated  PassengerEventType = "PassengerCreated"
	PassengerApproved PassengerEventType = "PassengerApproved"
)

type PassengerEvent struct {
	Type PassengerEventType `json:"type"`
	Data interface{}        `json:"data"`
}
