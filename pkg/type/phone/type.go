package phone

import "errors"

var ErrorIncorrectPhoneNumber = errors.New("incorrect phone number")

type PhoneNumber string

func New(phone string) (*PhoneNumber, error) {
	digits := make([]byte, 0, len(phone))
	for _, digit := range phone {
		if digit >= '0' && digit <= '9' {
			digits = append(digits, byte(digit))
		}
	}

	if len(digits) == 0 {
		return nil, ErrorIncorrectPhoneNumber
	}

	pn := PhoneNumber(digits)
	return &pn, nil
}
