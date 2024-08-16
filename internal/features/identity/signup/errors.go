package signup

import "fmt"

func duplicateUsername(username string) error {
	return fmt.Errorf("ERR_004: Username/Phone number is already registered: %s", username)
}
