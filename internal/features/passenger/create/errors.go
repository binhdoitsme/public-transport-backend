package create

import "fmt"

func alreadyExistsError(phoneNumber string, vneID string) error {
	return fmt.Errorf("ERR_012: Phone number (%s) and/or VNEID (%s) already registered", phoneNumber, vneID)
}
