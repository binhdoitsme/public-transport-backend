package errors

import "fmt"

func ToGenericError(err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("ERR_000: %s", err.Error())
}

func ToValidationError(err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("ERR_001: %s", err.Error())
}

func NotAnAdminError() error {
	return fmt.Errorf("ERR_002: Current user does not have privileges to perform this action")
}

func NotAuthorizedError() error {
	return fmt.Errorf("ERR_003: You need to log in to perform this action")
}
