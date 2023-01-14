package main

import (
	"github.com/cr00z/goContacts/pkg/store/postgres"
	http_delivery "github.com/cr00z/goContacts/services/contact/internal/delivery/http"
	repository "github.com/cr00z/goContacts/services/contact/internal/repository/postgres"
	contact_service "github.com/cr00z/goContacts/services/contact/internal/service/contact"
	group_service "github.com/cr00z/goContacts/services/contact/internal/service/group"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	conn, err := postgres.New(postgres.Settings{})
	if err != nil {
		panic(err)
	}
	defer conn.Pool.Close()
	var (
		repo       = repository.New(conn, repository.Options{})
		contactSvc = contact_service.NewContactService(repo, contact_service.Options{})
		groupSvc   = group_service.NewGroupService(repo, group_service.Options{})
		httpDeliv  = http_delivery.New(contactSvc, groupSvc, http_delivery.Options{})
	)

	go func() {
		// TODO change logger
		log.Print("service started on http port ", viper.GetUint("HTTP_PORT"))
		if err = httpDeliv.Run(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	// TODO change logger
	log.Print("signal received: ", <-quit, ", server shutting down")
}

// PGPORT=5442;PGDATABASE=test;PGUSER=user;PGPASSWORD=passwd
