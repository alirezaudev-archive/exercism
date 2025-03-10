package thefarm

import (
	"errors"
	"fmt"
)

type InvalidCowsError struct {
	cows    int
	message string
}

func (err *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", err.cows, err.message)
}

func DivideFood(calculator FodderCalculator, cows int) (float64, error) {
	amount, err := calculator.FodderAmount(cows)
	if err != nil {
		return 0, err
	}

	factor, err2 := calculator.FatteningFactor()
	if err2 != nil {
		return 0, err2
	}

	return (amount * factor) / float64(cows), nil
}

func ValidateInputAndDivideFood(calculator FodderCalculator, cows int) (float64, error) {
	if cows > 0 {
		return DivideFood(calculator, cows)
	}

	return 0, errors.New("invalid number of cows")
}

func ValidateNumberOfCows(cows int) error {
	if cows > 0 {
		return nil
	}

	err := &InvalidCowsError{cows: cows, message: "there are no negative cows"}
	if cows == 0 {
		err.message = "no cows don't need food"
	}

	return err
}
