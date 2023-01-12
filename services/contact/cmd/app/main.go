package main

import (
	"fmt"
	"github.com/cr00z/goContacts/pkg/store/postgres"
	"github.com/cr00z/goContacts/services/contact/internal/delivery"
	repository "github.com/cr00z/goContacts/services/contact/internal/repository/postgres"
	service "github.com/cr00z/goContacts/services/contact/internal/service"
)

func main() {
	conn, err := postgres.New(postgres.Settings{})
	if err != nil {
		panic(err)
	}
	defer conn.Pool.Close()

	repo := repository.New(conn)
	svc := service.New(repo)
	deliv := delivery.New(svc)

	fmt.Println("hello world", deliv)
}
