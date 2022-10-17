package shelly

import "fmt"

type InvalidShellyType string

func (e InvalidShellyType) Error() string {
	return fmt.Sprintf("Invalid Shelly type: %s", string(e))
}
