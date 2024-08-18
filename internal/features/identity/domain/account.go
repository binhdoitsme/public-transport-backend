package domain

import (
	"time"

	"github.com/bwmarrin/snowflake"
)

type Role string

const SuperAdmin Role = "SuperAdmin"
const Admin Role = "Admin"
const User Role = "User"

const DefaultPersonalImage string = "https://upload.wikimedia.org/wikipedia/commons/b/bc/Unknown_person.jpg"

const MinPasswordLength = 6
const RefreshTokenTTL = 7 * 24 * time.Hour

type RefreshToken struct {
	Token      string
	Expiration time.Time
}

type Account struct {
	Id            uint64
	Username      string
	Password      string
	Name          string
	Role          Role
	PersonalImage string
	RefreshTokens []RefreshToken
}

func New(
	username string,
	password string,
	name string,
	role Role,
	personalImage *string,
	id *uint64,
) (*Account, error) {
	if id == nil {
		node, err := snowflake.NewNode(16)
		if err != nil {
			return nil, err
		}
		generated := uint64(node.Generate().Int64())
		id = &generated
	}
	account := &Account{
		Id:            *id,
		Username:      username,
		Password:      password,
		Name:          name,
		Role:          role,
		RefreshTokens: make([]RefreshToken, 0),
	}
	if personalImage == nil {
		account.PersonalImage = DefaultPersonalImage
	} else {
		account.PersonalImage = *personalImage
	}
	return account, nil
}

func (account *Account) AddRefreshToken(token string) {
	if len(token) == 0 {
		return
	}
	account.RefreshTokens = append(account.RefreshTokens, RefreshToken{
		Token:      token,
		Expiration: time.Now().UTC().Add(RefreshTokenTTL),
	})
}

// InvalidateAllTokens removes all refresh tokens, useful when the password is changed.
func (account *Account) InvalidateAllTokens() {
	account.RefreshTokens = []RefreshToken{}
}

// RemoveExpiredTokens cleans up expired tokens from the account.
func (account *Account) RemoveExpiredTokens() {
	validTokens := []RefreshToken{}
	now := time.Now().UTC()
	for _, rt := range account.RefreshTokens {
		if rt.Expiration.After(now) {
			validTokens = append(validTokens, rt)
		}
	}
	account.RefreshTokens = validTokens
}

// InvalidateToken invalidates a specific refresh token by removing it from the account's RefreshTokens slice.
func (account *Account) InvalidateToken(token string) {
	if len(token) == 0 {
		return
	}
	// Iterate over the refresh tokens and filter out the one that matches the token to be invalidated.
	validTokens := []RefreshToken{}
	for _, rt := range account.RefreshTokens {
		if rt.Token != token {
			validTokens = append(validTokens, rt)
		}
	}
	account.RefreshTokens = validTokens
}
