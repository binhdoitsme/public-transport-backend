package support

import "context"

type PasswordServices interface {
	ToStoredForm(ctx context.Context, password string) (string, error)
}
