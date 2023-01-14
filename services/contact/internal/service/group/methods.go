package group

import (
	"github.com/cr00z/goContacts/pkg/type/queryparam"
	"github.com/cr00z/goContacts/services/contact/internal/domain/group"
	"github.com/google/uuid"
)

// Group interface

func (s *Service) Create(group *group.Group) error {
	return s.repo.CreateGroup(group)
}

func (s *Service) Update(group *group.Group) error {
	return s.repo.UpdateGroup(group)
}

func (s *Service) DeleteById(id uuid.UUID) error {
	return s.repo.DeleteGroupById(id)
}

// GroupReader interface

func (s *Service) Find(filter queryparam.QueryParam) ([]*group.Group, error) {
	return s.repo.FindGroups(filter)
}

func (s *Service) GetById(id uuid.UUID) (*group.Group, error) {
	return s.repo.GetGroupById(id)
}

func (s *Service) Count() error {
	return s.repo.CountGroups()
}

// ContactInGroup interface

func (s *Service) AddContactsToGroup(groupId uuid.UUID, contactIds []uuid.UUID) error {
	return s.repo.AddContactsToGroup(groupId, contactIds)
}

func (s *Service) RemoveContactsFromGroup(groupId uuid.UUID, contactIds []uuid.UUID) error {
	return s.repo.RemoveContactsFromGroup(groupId, contactIds)
}
