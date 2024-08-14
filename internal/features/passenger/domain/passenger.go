package domain

import (
	"fmt"
	"public-transport-backend/internal/common/dates"
	"time"

	"github.com/bwmarrin/snowflake"
)

type AccountType string

const (
	Individual AccountType = "Individual"
	Student    AccountType = "Student"
	Group      AccountType = "Group"
	Elder      AccountType = "Elder"
)

type AccountStatus string

const (
	PendingApproval AccountStatus = "PENDING_APPROVAL"
	Approved        AccountStatus = "APPROVED"
	Rejected        AccountStatus = "REJECTED"
	Archived        AccountStatus = "ARCHIVED" // pending deletion
)

const MinAgeForElderPass int = 60

type Account struct {
	Id                   uint64
	PhoneNumber          string
	VneID                string
	Name                 string
	DOB                  time.Time
	Gender               string
	PersonalImage        string
	AccountType          AccountType
	ConfirmationDocument *string
	Status               AccountStatus
}

func NewAccount(
	phoneNumber string,
	vneID string,
	name string,
	dob time.Time,
	gender string,
	personalImage string,
	accountType AccountType,
	confirmationDocument *string,
	id *uint64,
	status *AccountStatus,
) (*Account, error) {
	if id == nil {
		node, err := snowflake.NewNode(16)
		if err != nil {
			return nil, err
		}
		generated := uint64(node.Generate().Int64())
		id = &generated
	}
	if status == nil {
		defaultStatus := PendingApproval
		status = &defaultStatus
	}

	switch accountType {
	case Student:
	case Group:
		if confirmationDocument == nil {
			return nil, fmt.Errorf("ERR_010: Confirmation document is required for passenger type %s", accountType)
		}
	case Elder:
		if dob.AddDate(MinAgeForElderPass, 0, 0).Before(dates.StartOfDay(time.Now())) {
			return nil, fmt.Errorf("ERR_011: Must be over %d years old to be eligible for Elder passenger type", MinAgeForElderPass)
		}
	}

	return &Account{
		Id:                   *id,
		PhoneNumber:          phoneNumber,
		VneID:                vneID,
		Name:                 name,
		DOB:                  dob,
		Gender:               gender,
		PersonalImage:        personalImage,
		AccountType:          accountType,
		ConfirmationDocument: confirmationDocument,
		Status:               *status,
	}, nil
}
