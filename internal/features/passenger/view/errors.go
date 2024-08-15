package view

import "fmt"

func notFound(id uint64) error {
	return fmt.Errorf("ERR_013: Passenger not found for ID %d", id)
}
