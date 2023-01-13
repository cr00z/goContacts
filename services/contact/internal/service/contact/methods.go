package contact

import (
	"github.com/cr00z/goContacts/pkg/type/queryparam"
	"github.com/cr00z/goContacts/services/contact/internal/domain/contact"
	"github.com/google/uuid"
)

// Contact interface

func (s Service) Create(contact *contact.Contact) error {
	panic("need implementation")
}

func (s Service) Update(contact *contact.Contact) error {
	panic("need implementation")
}

func (s Service) DeleteById(id uuid.UUID) error {
	panic("need implementation")
}

// ContactReader interface

func (s Service) Find(filter queryparam.QueryParam) ([]*contact.Contact, error) {
	panic("need implementation")
}

func (s Service) GetById(id uuid.UUID) (*contact.Contact, error) {
	panic("need implementation")
}

func (s Service) Count() error {
	panic("need implementation")
}
