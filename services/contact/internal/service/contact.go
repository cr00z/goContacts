package service

import "github.com/cr00z/goContacts/services/contact/internal/service/adapters/storage"

type ContactService struct {
	repo storage.Storager
}

func New(repo storage.Storager) Servicer {
	return &ContactService{repo}
}
