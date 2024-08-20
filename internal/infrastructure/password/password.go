package password

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type PasswordServices struct{}

func (p *PasswordServices) ToStoredForm(ctx context.Context, password string) (string, error) {
	if password == "" {
		return "", errors.New("password cannot be empty")
	}
	pwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(pwd), err
}

func (p *PasswordServices) Compare(ctx context.Context, stored string, entered string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(stored), []byte(entered))
	return err != nil
}
