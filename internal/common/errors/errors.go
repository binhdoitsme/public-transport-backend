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

func NotAnAdminError(userId uint64) error {
	return fmt.Errorf("ERR_002: UserId %d does not have privileges to perform this action", userId)
}
