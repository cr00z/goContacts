package group

import (
	"github.com/cr00z/opsiot-cnts-demo/services/contact/internal/service/adapters/storage"
)

type Options struct{}

type Service struct {
	repo    storage.Storager
	options Options
}

func NewGroupService(repo storage.Storager, opts Options) *Service {
	s := &Service{
		repo: repo,
	}
	s.SetOptions(opts)
	return s
}

func (s *Service) SetOptions(opts Options) {
	if s.options != opts {
		s.options = opts
	}
}
