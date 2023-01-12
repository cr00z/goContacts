package repository

import (
	"github.com/cr00z/goContacts/pkg/store/postgres"
	"github.com/cr00z/goContacts/services/contact/internal/repository"
)

type PostgresRepository struct {
	Store *postgres.Store
}

func New(store *postgres.Store) repository.Storager {
	return &PostgresRepository{store}
}
