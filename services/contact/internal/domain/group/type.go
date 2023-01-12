package group

import (
	"github.com/google/uuid"
	"time"
)

type Group struct {
	ID           uuid.UUID
	CreatedAt    time.Time
	ModifiedAt   time.Time
	Name         string
	Description  string
	ContactCount int64
	IsArchived   bool
}

func New() Group {
	return Group{
		ID:         uuid.New(),
		CreatedAt:  time.Now(),
		ModifiedAt: time.Now(),
	}
}

func (g *Group) SetName(name string) error {
	if len(name) > 250 {
		return ErrorGroupNameIsLong
	}
	g.Name = name
	return nil
}
