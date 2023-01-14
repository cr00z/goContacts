package repository

import (
	"github.com/cr00z/goContacts/pkg/store/postgres"
	"github.com/cr00z/goContacts/pkg/type/queryparam"
	"github.com/cr00z/goContacts/services/contact/internal/domain/contact"
	"github.com/cr00z/goContacts/services/contact/internal/domain/group"
	"github.com/google/uuid"
)

type Options struct {
}

type PostgresRepository struct {
	Store   *postgres.Store
	options Options
}

func New(store *postgres.Store, opts Options) *PostgresRepository {
	r := &PostgresRepository{
		Store: store,
	}
	r.SetOptions(opts)
	return r
}

func (r *PostgresRepository) SetOptions(opts Options) {
	if r.options != opts {
		r.options = opts
	}
}

// TODO: растащить по файлам
// Contact interface

func (r *PostgresRepository) CreateContact(contact *contact.Contact) error {
	return nil
}

func (r *PostgresRepository) UpdateContact(
	contactId uuid.UUID,
	updateFn func(oldContact *contact.Contact) (*contact.Contact, error),
) error {
	return nil
}

func (r *PostgresRepository) DeleteContactById(id uuid.UUID) error {
	return nil
}

// ContactReader interface

func (r *PostgresRepository) FindContacts(filter queryparam.QueryParam) ([]*contact.Contact, error) {
	return nil, nil
}

func (r *PostgresRepository) GetContactById(id uuid.UUID) (*contact.Contact, error) {
	return nil, nil
}

func (r *PostgresRepository) CountContacts() error {
	return nil
}

// Group interface

func (r *PostgresRepository) CreateGroup(group *group.Group) error {
	return nil
}

func (r *PostgresRepository) UpdateGroup(group *group.Group) error {
	return nil
}

func (r *PostgresRepository) DeleteGroupById(id uuid.UUID) error {
	return nil
}

// GroupReader interface
func (r *PostgresRepository) FindGroups(filter queryparam.QueryParam) ([]*group.Group, error) {
	return nil, nil
}

func (r *PostgresRepository) GetGroupById(id uuid.UUID) (*group.Group, error) {
	return nil, nil
}

func (r *PostgresRepository) CountGroups() error {
	return nil
}

// ContactInGroup interface

func (r *PostgresRepository) AddContactsToGroup(groupId uuid.UUID, contactIds []uuid.UUID) error {
	return nil
}

func (r *PostgresRepository) RemoveContactsFromGroup(groupId uuid.UUID, contactIds []uuid.UUID) error {
	return nil
}
