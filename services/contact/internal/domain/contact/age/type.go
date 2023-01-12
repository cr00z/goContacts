package age

import (
	"github.com/pkg/errors"
	"strconv"
)

var (
	ageMaxLength      = 200
	ErrorAgeIncorrect = errors.Errorf("Age incorrect, must be less or equal %d", ageMaxLength)
)

type Age uint8

func New(age int) (Age, error) {
	if age < 0 || age > ageMaxLength {
		return Age(0), ErrorAgeIncorrect
	}
	return Age(age), nil
}

func (a Age) String() string {
	return strconv.Itoa(int(a))
}
