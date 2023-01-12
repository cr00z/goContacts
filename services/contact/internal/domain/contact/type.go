package contact

import (
	"github.com/google/uuid"
	"time"
)

type Contact struct {
	ID          uuid.UUID
	CreatedAt   time.Time
	ModifiedAt  time.Time
	Name        string
	Surname     string
	Patronymic  string
	Email       string
	PhoneNumber string
	Age         int16
	Gender      int16
}

func New() Contact {
	return Contact{
		ID:         uuid.New(),
		CreatedAt:  time.Now(),
		ModifiedAt: time.Now(),
	}
}

func (c *Contact) SetPhoneNumber(number string) error {
	for _, digit := range number {
		if digit < '0' || digit > '9' {
			return ErrorIncorrectPhoneNumber
		}
	}
	c.PhoneNumber = number
	return nil
}
