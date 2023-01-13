package service

import (
	"github.com/cr00z/goContacts/services/contact/internal/domain/contact"
	"github.com/google/uuid"
)

type Servicer interface {
}

type Contact interface {
	Create(contact contact.Contact) error
	Update(contact contact.Contact) error
	Delete(id uuid.UUID) error
}

type ContactReader interface {
	List(param queryparam.queryParam)
}
