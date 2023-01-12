package description

import (
	"github.com/pkg/errors"
	"unicode/utf8"
)

var (
	descriptionMaxLength      = 1000
	ErrorDescriptionIncorrect = errors.Errorf(
		"Description len incorrect, must be less or equal %d", descriptionMaxLength)
)

type Description string

func New(desc string) (*Description, error) {
	if utf8.RuneCountInString(desc) > descriptionMaxLength {
		return nil, ErrorDescriptionIncorrect
	}
	d := Description(desc)
	return &d, nil
}

func (d Description) String() string {
	return string(d)
}
