// Package thefarm does some cow fodder calculations.
package thefarm

import (
	"errors"
	"fmt"
)

type InvalidCowsError struct {
	cows int
	msg  string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.cows, e.msg)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(calculator FodderCalculator, cows int) (float64, error) {
	total, err := calculator.FodderAmount(cows)
	if err != nil {
		return 0, errDeterminingAmount
	}
	fattening, err := calculator.FatteningFactor()
	if err != nil {
		return 0, errDeterminingFactor
	}
	return total * fattening / float64(cows), nil
}

func ValidateInputAndDivideFood(calculator FodderCalculator, cows int) (float64, error) {
	if cows <= 0 {
		return 0, errors.New("invalid number of cows")
	}
	return DivideFood(calculator, cows)
}

func ValidateNumberOfCows(cows int) *InvalidCowsError {
	switch {
	case cows < 0:
		return &InvalidCowsError{cows, "there are no negative cows"}
	case cows == 0:
		return &InvalidCowsError{cows, "no cows don't need food"}
	default:
		return nil
	}
}
