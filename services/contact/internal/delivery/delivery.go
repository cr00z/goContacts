package delivery

import "github.com/cr00z/goContacts/services/contact/internal/service"

type ContactDelivery struct {
	Service service.Servicer
}

func New(svc service.Servicer) *ContactDelivery {
	return &ContactDelivery{svc}
}
