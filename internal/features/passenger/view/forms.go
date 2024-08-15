package view

type RequestingUser struct {
	UserId uint64
}

type PassengerByIdForm struct {
	Id uint64
}

type AdminPassengerByIdForm struct {
	Id uint64
	*RequestingUser
}
