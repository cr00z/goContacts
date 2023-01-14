package storage

import (
	"github.com/cr00z/goContacts/pkg/type/queryparam"
	"github.com/cr00z/goContacts/services/contact/internal/domain/contact"
	"github.com/cr00z/goContacts/services/contact/internal/domain/group"
	"github.com/google/uuid"
)

type Storager interface {
	Contact
	ContactReader
	Group
	GroupReader
	ContactInGroup
}

type Contact interface {
	CreateContact(contact *contact.Contact) error
	UpdateContact(
		contactId uuid.UUID,
		updateFn func(oldContact *contact.Contact) (*contact.Contact, error),
	) error
	DeleteContactById(id uuid.UUID) error
}

type ContactReader interface {
	FindContacts(filter queryparam.QueryParam) ([]*contact.Contact, error)
	GetContactById(id uuid.UUID) (*contact.Contact, error)
	CountContacts() error
}

type Group interface {
	CreateGroup(group *group.Group) error
	UpdateGroup(group *group.Group) error
	DeleteGroupById(id uuid.UUID) error
}

type GroupReader interface {
	FindGroups(filter queryparam.QueryParam) ([]*group.Group, error)
	GetGroupById(id uuid.UUID) (*group.Group, error)
	CountGroups() error
}

type ContactInGroup interface {
	AddContactsToGroup(groupId uuid.UUID, contactIds []uuid.UUID) error
	RemoveContactsFromGroup(groupId uuid.UUID, contactIds []uuid.UUID) error
}
