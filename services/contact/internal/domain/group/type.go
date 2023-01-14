package group

import (
	"github.com/cr00z/opsiot-cnts-demo/services/contact/internal/domain/group/description"
	"github.com/cr00z/opsiot-cnts-demo/services/contact/internal/domain/group/name"
	"github.com/google/uuid"
	"time"
)

type Group struct {
	id           uuid.UUID
	createdAt    time.Time
	modifiedAt   time.Time
	name         name.Name
	description  description.Description
	contactCount int64
	isArchived   bool
}

func New(
	name name.Name,
	description description.Description,
	contactCount int64,
	isArchived bool,
) *Group {
	id := uuid.New()
	tm := time.Now().UTC()
	return NewWithId(id, tm, tm, name, description, contactCount, isArchived)
}

func NewWithId(
	ID uuid.UUID,
	createdAt time.Time,
	modifiedAt time.Time,
	name name.Name,
	description description.Description,
	contactCount int64,
	isArchived bool,
) *Group {
	return &Group{
		id:           ID,
		createdAt:    createdAt,
		modifiedAt:   modifiedAt,
		name:         name,
		description:  description,
		contactCount: contactCount,
		isArchived:   isArchived,
	}
}

func (g Group) Id() uuid.UUID {
	return g.id
}

func (g Group) CreatedAt() time.Time {
	return g.createdAt
}

func (g Group) ModifiedAt() time.Time {
	return g.modifiedAt
}

func (g Group) Name() name.Name {
	return g.name
}

func (g Group) Description() description.Description {
	return g.description
}

func (g Group) ContactCount() int64 {
	return g.contactCount
}

func (g Group) IsArchived() bool {
	return g.isArchived
}
