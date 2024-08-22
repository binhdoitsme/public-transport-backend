package models

import (
	passenger "public-transport-backend/internal/features/passenger/domain"
	"time"

	"gorm.io/gorm"
)

type Passenger struct {
	gorm.Model
	ID                   uint64 `gorm:"primaryKey;autoIncrement:false"`
	PhoneNumber          string `gorm:"uniqueIndex:idx_unique_passenger;size:16"`
	VneID                string `gorm:"uniqueIndex:idx_unique_passenger;size:16"`
	Name                 string
	DOB                  time.Time
	Gender               string
	PersonalImage        string
	AccountType          string
	ConfirmationDocument *string
	Status               string
}

func (p *Passenger) ToPassenger() *passenger.Account {
	return &passenger.Account{
		Id:                   p.ID,
		PhoneNumber:          p.PhoneNumber,
		VneID:                p.VneID,
		Name:                 p.Name,
		DOB:                  p.DOB,
		Gender:               p.Gender,
		PersonalImage:        p.PersonalImage,
		AccountType:          passenger.AccountType(p.AccountType),
		ConfirmationDocument: p.ConfirmationDocument,
		Status:               passenger.AccountStatus(p.Status),
	}
}
