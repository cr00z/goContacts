package email

import (
	"errors"
	"regexp"
	"strings"
)

var (
	emailRegexp       = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	ErrorEmailInvalid = errors.New("invalid email format")
)

type Email string

func New(email string) (*Email, error) {
	if !emailRegexp.MatchString(strings.ToLower(email)) {
		return nil, ErrorEmailInvalid
	}
	m := Email(email)
	return &m, nil
}

func (e Email) String() string {
	return string(e)
}
