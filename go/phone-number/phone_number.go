package phonenumber

import (
	"errors"
	"fmt"
	"regexp"
)

var (
	errInvalid = errors.New("invalid phone number")
	phoneRe    = regexp.MustCompile(`^[2-9][0-9]{2}[2-9][0-9]{6}$`)
)

func Number(phoneNumber string) (string, error) {
	var cleanNumber []rune
	for _, r := range phoneNumber {
		if r >= '0' && r <= '9' {
			cleanNumber = append(cleanNumber, r)
		}
	}
	if len(cleanNumber) == 11 {
		if cleanNumber[0] != '1' {
			return "", errInvalid
		}
		cleanNumber = cleanNumber[1:]
	}
	strCleanNumber := string(cleanNumber)
	if phoneRe.MatchString(strCleanNumber) {
		return strCleanNumber, nil
	}
	return "", errInvalid
}

func AreaCode(phoneNumber string) (string, error) {
	cleanNumber, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return cleanNumber[0:3], nil
}

func Format(phoneNumber string) (string, error) {
	cleanNumber, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(
		"(%s) %s-%s",
		cleanNumber[0:3], cleanNumber[3:6], cleanNumber[6:],
	), nil
}
