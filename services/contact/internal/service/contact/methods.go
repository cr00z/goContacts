package contact

import (
	"github.com/cr00z/goContacts/pkg/type/queryparam"
	"github.com/cr00z/goContacts/services/contact/internal/domain/contact"
	"github.com/google/uuid"
	"time"
)

// Contact interface

func (s *Service) Create(contact *contact.Contact) error {
	return s.repo.CreateContact(contact)
}

func (s *Service) Update(updContact *contact.Contact) error {
	return s.repo.UpdateContact(updContact.Id(), func(oldContact *contact.Contact) (*contact.Contact, error) {
		return contact.NewWithId(
			oldContact.Id(),
			oldContact.CreatedAt(),
			time.Now().UTC(),
			updContact.Name(),
			updContact.Surname(),
			updContact.Patronymic(),
			updContact.Email(),
			updContact.PhoneNumber(),
			updContact.Age(),
			updContact.Gender(),
		)
	})
}

func (s *Service) DeleteById(id uuid.UUID) error {
	return s.repo.DeleteContactById(id)
}

// ContactReader interface

func (s *Service) Find(filter queryparam.QueryParam) ([]*contact.Contact, error) {
	return s.repo.FindContacts(filter)
}

func (s *Service) GetById(id uuid.UUID) (*contact.Contact, error) {
	return s.repo.GetContactById(id)
}

func (s *Service) Count() error {
	return s.repo.CountContacts()
}
