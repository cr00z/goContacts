package http

import (
	"fmt"
	"github.com/cr00z/goContacts/services/contact/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault("HTTP_HOST", "localhost")
	viper.SetDefault("HTTP_PORT", 8080)
	viper.SetDefault("IS_PROD", "false")
}

type Options struct {
}

type ContactDelivery struct {
	contactSvc service.Contacter
	groupSvc   service.Grouper
	router     *gin.Engine
	options    Options
}

func New(cs service.Contacter, gs service.Grouper, opts Options) *ContactDelivery {
	d := &ContactDelivery{
		contactSvc: cs,
		groupSvc:   gs,
	}

	d.SetOptions(opts)
	d.router = d.InitRoutes()

	return d
}

func (d *ContactDelivery) SetOptions(opts Options) {
	if d.options != opts {
		d.options = opts
		// + dynamic refresh
	}
}

func (d *ContactDelivery) Run() error {
	return d.router.Run(fmt.Sprintf("%s:%d", viper.Get("HTTP_HOST"), viper.Get("HTTP_PORT")))
}
