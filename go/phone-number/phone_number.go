package phonenumber

import (
	"errors"
	"fmt"
	"regexp"
)

var nonDigitRegexp = regexp.MustCompile(`[^0-9]`)

func Number(phoneNumber string) (string, error) {
	digits := nonDigitRegexp.ReplaceAllString(phoneNumber, "")

	if len(digits) == 11 && digits[0] == '1' {
		digits = digits[1:]
	} else if len(digits) != 10 {
		return "", errors.New("")
	}

	if digits[0] < '2' || digits[3] < '2' {
		return "", errors.New("")
	}

	return digits, nil
}

func AreaCode(phoneNumber string) (string, error) {
	num, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return num[:3], nil
}

func Format(phoneNumber string) (string, error) {
	num, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("(%s) %s-%s", num[:3], num[3:6], num[6:]), nil
}
