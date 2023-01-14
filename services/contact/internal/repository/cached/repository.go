package cached

import (
	"github.com/cr00z/goContacts/pkg/type/queryparam"
	"github.com/cr00z/goContacts/services/contact/internal/domain/contact"
	"github.com/cr00z/goContacts/services/contact/internal/domain/group"
	"github.com/cr00z/goContacts/services/contact/internal/service/adapters/storage"
	"github.com/google/uuid"
)

type Options struct {
}

type Repository struct {
	options Options
	cache   map[string]interface{}
}

func New(storage storage.Storager, opts Options) *Repository {
	r := &Repository{}
	return r
}

func (r Repository) CreateContact(contact *contact.Contact) error {
	err := r.CreateContact(contact)
	// cache me

	//TODO implement me
	panic("implement me")
	return err
}

func (r Repository) UpdateContact(
	contactId uuid.UUID,
	updateFn func(oldContact *contact.Contact) (*contact.Contact, error),
) error {
	//TODO implement me
	panic("implement me")
}

func (r Repository) DeleteContactById(id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (r Repository) FindContacts(filter queryparam.QueryParam) ([]*contact.Contact, error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) GetContactById(id uuid.UUID) (*contact.Contact, error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) CountContacts() error {
	//TODO implement me
	panic("implement me")
}

func (r Repository) CreateGroup(group *group.Group) error {
	//TODO implement me
	panic("implement me")
}

func (r Repository) UpdateGroup(group *group.Group) error {
	//TODO implement me
	panic("implement me")
}

func (r Repository) DeleteGroupById(id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (r Repository) FindGroups(filter queryparam.QueryParam) ([]*group.Group, error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) GetGroupById(id uuid.UUID) (*group.Group, error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) CountGroups() error {
	//TODO implement me
	panic("implement me")
}

func (r Repository) AddContactsToGroup(groupId uuid.UUID, contactIds []uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (r Repository) RemoveContactsFromGroup(groupId uuid.UUID, contactIds []uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}
