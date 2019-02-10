package card

import (
	"errors"
	"strconv"
)

// ByID func
func ByID(index uint8) (id ID, err error) {
	if index < 1 {
		err = errors.New("Index cannot be less than 1")
	} else if index > 40 {
		err = errors.New("Index cannot be more than 40")
	} else {
		id = ID(index)
	}
	return
}

