package repository

import (
	"context"
	"github.com/cr00z/opsiot-cnts-demo/services/contact/internal/domain/contact"
	"github.com/google/uuid"
)

// Contact interface

func (r *PostgresRepository) CreateContact(contact *contact.Contact) error {
	const query = `INSERT INTO contact VALUES (
	      $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11                      
	);`
	_, err := r.Store.Pool.Exec(context.TODO(), query,
		contact.Id(),
		contact.CreatedAt(),
		contact.ModifiedAt(),
		contact.Name(),
		contact.Surname(),
		contact.Patronymic(),
		contact.Email(),
		contact.PhoneNumber(),
		contact.Age(),
		contact.Gender(),
		false,
	)
	return err
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
