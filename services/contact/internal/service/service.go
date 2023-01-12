package service

import "github.com/cr00z/goContacts/services/contact/internal/repository"

type Servicer interface {
}

type ContactService struct {
	repo repository.Storager
}

func New(repo repository.Storager) Servicer {
	return &ContactService{repo}
}
