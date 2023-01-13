package service

import (
	"github.com/cr00z/goContacts/pkg/type/queryparam"
	"github.com/cr00z/goContacts/services/contact/internal/domain/contact"
	"github.com/cr00z/goContacts/services/contact/internal/domain/group"
	"github.com/google/uuid"
)

type Contacter interface {
	Create(contact *contact.Contact) error
	Update(contact *contact.Contact) error
	DeleteById(id uuid.UUID) error
	ContactReader
}

type ContactReader interface {
	Find(filter queryparam.QueryParam) ([]*contact.Contact, error)
	GetById(id uuid.UUID) (*contact.Contact, error)
	Count() error
}

type Grouper interface {
	Create(group *group.Group) error
	Update(group *group.Group) error
	DeleteById(id uuid.UUID) error
	GroupReader
	ContactInGroup
}

type GroupReader interface {
	Find(filter queryparam.QueryParam) ([]*group.Group, error)
	GetById(id uuid.UUID) (*group.Group, error)
	Count() error
}

type ContactInGroup interface {
	AddContactsToGroup(groupId uuid.UUID, contactIds []uuid.UUID) error
	RemoveContactsFromGroup(groupId uuid.UUID, contactIds []uuid.UUID) error
}
