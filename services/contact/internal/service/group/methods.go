package group

import (
	"github.com/cr00z/goContacts/pkg/type/queryparam"
	"github.com/cr00z/goContacts/services/contact/internal/domain/group"
	"github.com/google/uuid"
)

// Group interface

func (s Service) Create(group *group.Group) error {
	panic("need implementation")
}

func (s Service) Update(group *group.Group) error {
	panic("need implementation")
}

func (s Service) DeleteById(id uuid.UUID) error {
	panic("need implementation")
}

// GroupReader interface

func (s Service) Find(filter queryparam.QueryParam) ([]*group.Group, error) {
	panic("need implementation")
}

func (s Service) GetById(id uuid.UUID) (*group.Group, error) {
	panic("need implementation")
}

func (s Service) Count() error {
	panic("need implementation")
}

// ContactInGroup interface

func (s Service) AddContactsToGroup(groupId uuid.UUID, contactIds []uuid.UUID) error {
	panic("need implementation")
}

func (s Service) RemoveContactsFromGroup(groupId uuid.UUID, contactIds []uuid.UUID) error {
	panic("need implementation")
}
