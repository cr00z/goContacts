package contact

import (
	"github.com/cr00z/opsiot-cnts-demo/pkg/type/email"
	"github.com/cr00z/opsiot-cnts-demo/pkg/type/gender"
	"github.com/cr00z/opsiot-cnts-demo/pkg/type/phone"
	"github.com/cr00z/opsiot-cnts-demo/services/contact/internal/domain/contact/age"
	"github.com/cr00z/opsiot-cnts-demo/services/contact/internal/domain/contact/name"
	"github.com/cr00z/opsiot-cnts-demo/services/contact/internal/domain/contact/patronymic"
	"github.com/cr00z/opsiot-cnts-demo/services/contact/internal/domain/contact/surname"
	"github.com/google/uuid"
	"time"
)

type Contact struct {
	id         uuid.UUID
	createdAt  time.Time
	modifiedAt time.Time

	name       name.Name
	surname    surname.Surname
	patronymic patronymic.Patronymic

	email       email.Email
	phoneNumber phone.PhoneNumber

	age    age.Age
	gender gender.Gender
}

func New(
	name name.Name,
	surname surname.Surname,
	patronymic patronymic.Patronymic,
	email email.Email,
	phoneNumber phone.PhoneNumber,
	age age.Age,
	gender gender.Gender,
) (*Contact, error) {
	id := uuid.New()
	tm := time.Now().UTC()
	return NewWithId(id, tm, tm, name, surname, patronymic, email, phoneNumber, age, gender)
}

func NewWithId(
	id uuid.UUID,
	createdAt time.Time,
	modifiedAt time.Time,
	name name.Name,
	surname surname.Surname,
	patronymic patronymic.Patronymic,
	email email.Email,
	phoneNumber phone.PhoneNumber,
	age age.Age,
	gender gender.Gender,
) (*Contact, error) {
	return &Contact{
		id:          id,
		createdAt:   createdAt,
		modifiedAt:  modifiedAt,
		name:        name,
		surname:     surname,
		patronymic:  patronymic,
		email:       email,
		phoneNumber: phoneNumber,
		age:         age,
		gender:      gender,
	}, nil
}

func (c Contact) Id() uuid.UUID {
	return c.id
}

func (c Contact) CreatedAt() time.Time {
	return c.createdAt
}

func (c Contact) ModifiedAt() time.Time {
	return c.modifiedAt
}

func (c Contact) Name() name.Name {
	return c.name
}

func (c Contact) Surname() surname.Surname {
	return c.surname
}

func (c Contact) Patronymic() patronymic.Patronymic {
	return c.patronymic
}

func (c Contact) Email() email.Email {
	return c.email
}

func (c Contact) PhoneNumber() phone.PhoneNumber {
	return c.phoneNumber
}

func (c Contact) Age() age.Age {
	return c.age
}

func (c Contact) Gender() gender.Gender {
	return c.gender
}
