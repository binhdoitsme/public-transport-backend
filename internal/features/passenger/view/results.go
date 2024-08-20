package view

import (
	passenger "public-transport-backend/internal/features/passenger/domain"
	"time"
)

type PassengerResult struct {
	Id                   uint64                  `json:"id"`
	PhoneNumber          string                  `json:"phoneNumber"`
	VneID                string                  `json:"vneId"`
	Name                 string                  `json:"name"`
	DOB                  time.Time               `json:"dob"`
	Gender               string                  `json:"gender"`
	PersonalImage        string                  `json:"personalImage"`
	AccountType          passenger.AccountType   `json:"accountType"`
	ConfirmationDocument *string                 `json:"confirmationDocument"`
	Status               passenger.AccountStatus `json:"status"`
}

func ToResult(p *passenger.Account) *PassengerResult {
	return &PassengerResult{
		Id:                   p.Id,
		PhoneNumber:          p.PhoneNumber,
		VneID:                p.VneID,
		Name:                 p.Name,
		DOB:                  p.DOB,
		Gender:               p.Gender,
		PersonalImage:        p.PersonalImage,
		AccountType:          p.AccountType,
		ConfirmationDocument: p.ConfirmationDocument,
		Status:               p.Status,
	}
}

func ToResults(ps []passenger.Account) []PassengerResult {
	results := make([]PassengerResult, 0, len(ps))
	for _, p := range ps {
		results = append(results, *ToResult(&p))
	}
	return results
}
