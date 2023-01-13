package main

import (
	"fmt"
	"github.com/cr00z/goContacts/pkg/store/postgres"
	"github.com/cr00z/goContacts/services/contact/internal/delivery/http"
	repository "github.com/cr00z/goContacts/services/contact/internal/repository/postgres"
	contact_service "github.com/cr00z/goContacts/services/contact/internal/service/contact"
	group_service "github.com/cr00z/goContacts/services/contact/internal/service/group"
)

func main() {
	conn, err := postgres.New(postgres.Settings{})
	if err != nil {
		panic(err)
	}
	defer conn.Pool.Close()

	repo := repository.New(conn, repository.Options{})
	contactSvc := contact_service.NewContactService(repo, contact_service.Options{})
	groupSvc := group_service.NewGroupService(repo, group_service.Options{})
	deliv := http.New(contactSvc, groupSvc, http.Options{})

	fmt.Println("hello world", deliv)
}

// PGPORT=5442;PGDATABASE=test;PGUSER=user;PGPASSWORD=passwd
