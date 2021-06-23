package helper

import (
	"errors"
	"math"
	"strconv"
)

func ValidateIDNumber(id string) error {
	if num, err := strconv.Atoi(id); err != nil || num == 0 || math.Signbit(float64(num)) == true {
		return errors.New("input must be a valid id")
	}

	return nil
}
