package models

import (
	identity "public-transport-backend/internal/features/identity/domain"
	"time"

	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	ID            uint64 `gorm:"primaryKey;autoIncrement:false"`
	Username      string
	Password      string `gorm:"->:false"`
	Name          string
	Role          string
	PersonalImage string
	RefreshTokens []RefreshToken
}

type RefreshToken struct {
	AccountID  uint64 `gorm:"primaryKey;autoIncrement:false"`
	Token      string `gorm:"index;primaryKey;autoIncrement:false"`
	Expiration time.Time
}

func (account *Account) ToAccount() *identity.Account {
	refreshTokens := make([]identity.RefreshToken, 0, len(account.RefreshTokens))
	for _, token := range account.RefreshTokens {
		refreshTokens = append(
			refreshTokens,
			identity.RefreshToken{
				Token:      token.Token,
				Expiration: token.Expiration,
			})
	}
	acc := &identity.Account{
		Id:            account.ID,
		Username:      account.Username,
		Name:          account.Name,
		Role:          identity.Role(account.Role),
		PersonalImage: account.PersonalImage,
		RefreshTokens: refreshTokens,
	}
	return acc
}
