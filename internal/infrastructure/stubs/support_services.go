package stubs

import (
	"context"
	"errors"
	"fmt"
	identity "public-transport-backend/internal/features/identity/domain"
	"strconv"
	"strings"
	"time"
)

// TokenServicesStub is a stub implementation of the TokenServices interface.
type TokenServicesStub struct {
	refreshTokens map[string]*identity.Account
	accessTokens  map[string]*identity.Account
}

// NewTokenServicesStub creates a new instance of TokenServicesStub with some seeded data.
func NewTokenServices() *TokenServicesStub {
	return &TokenServicesStub{
		refreshTokens: make(map[string]*identity.Account),
		accessTokens:  make(map[string]*identity.Account),
	}
}

// NewRefreshToken generates a new refresh token for a given account.
func (s *TokenServicesStub) NewRefreshToken(ctx context.Context, account *identity.Account) (string, error) {
	if account == nil {
		return "", errors.New("account cannot be nil")
	}
	// Simulate refresh token generation (a simple string with account ID and timestamp)
	token := fmt.Sprintf("refresh-%d-%d", account.Id, time.Now().UnixNano())
	// Store the token with the associated account
	s.refreshTokens[token] = account
	return token, nil
}

// NewAccessToken generates a new access token for a given account, based on a refresh token.
func (s *TokenServicesStub) NewAccessToken(ctx context.Context, account *identity.Account, refreshToken string) (string, error) {
	if account == nil {
		return "", errors.New("account cannot be nil")
	}

	// Check if the refresh token is valid
	if _, exists := s.refreshTokens[refreshToken]; !exists {
		return "", errors.New("invalid refresh token")
	}

	// Simulate access token generation (a simple string with account ID and timestamp)
	token := fmt.Sprintf("access-%d-%d", account.Id, time.Now().UnixNano())
	// Store the token with the associated account
	s.accessTokens[token] = account
	return token, nil
}

// Parse retrieves the account associated with a given token.
func (s *TokenServicesStub) Parse(token string) (*identity.Account, error) {
	// Example of token format: "access-<account_id>-<timestamp>" or "refresh-<account_id>-<timestamp>"
	split := strings.Split(token, "-")
	n := len(split)
	if n < 3 {
		return nil, errors.New("invalid token format")
	}
	accountId, err := strconv.Atoi(split[1])
	tokenType := split[0]
	if err != nil {
		return nil, errors.New("invalid token format")
	}
	ts, err := strconv.ParseInt(split[2], 10, 64)
	if err != nil {
		return nil, errors.New("invalid token format")
	}
	t := time.Unix(ts / 1_000_000_000, 0)
	now := time.Now().UTC()
	if now.Sub(t).Seconds() > 3600 {
		return nil, errors.New("expired token")
	}

	// Find the account associated with the extracted account ID
	account := &identity.Account{Id: uint64(accountId)}

	// Optionally, check if the token type is valid (e.g., "access" or "refresh")
	if tokenType != "access" && tokenType != "refresh" {
		return nil, errors.New("invalid token type")
	}

	// Return the account if found and valid
	return account, nil
}

// PasswordServicesStub is a stub implementation of the PasswordServices interface.
type PasswordServicesStub struct{}

// NewPasswordServices creates a new instance of PasswordServicesStub.
func NewPasswordServices() *PasswordServicesStub {
	return &PasswordServicesStub{}
}

// ToStoredForm converts a plain password to its stored form (hash).
func (p *PasswordServicesStub) ToStoredForm(ctx context.Context, password string) (string, error) {
	if password == "" {
		return "", errors.New("password cannot be empty")
	}
	// In a real implementation, this would hash the password, but here we just append "_hashed" for the stub.
	return password, nil
}

func (p *PasswordServicesStub) Compare(ctx context.Context, stored string, entered string) bool {
	return entered != stored
}
