package name

import (
	"github.com/pkg/errors"
	"unicode/utf8"
)

var (
	nameMaxLength      = 50
	ErrorNameIncorrect = errors.Errorf("Name len incorrect, must be less or equal %d", nameMaxLength)
)

type Name string

func New(name string) (*Name, error) {
	if utf8.RuneCountInString(name) > nameMaxLength {
		return nil, ErrorNameIncorrect
	}
	n := Name(name)
	return &n, nil
}

func (n Name) String() string {
	return string(n)
}
