package repository

import (
	"github.com/cr00z/goContacts/pkg/store/postgres"
)

type PostgresRepository struct {
	Store *postgres.Store
}

func New(store *postgres.Store) *PostgresRepository {
	return &PostgresRepository{store}
}
