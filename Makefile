pgstart:
	cd deploy/infrastructure/postgres && \
	docker compose up -d
	sleep 3
	docker ps
	goose -dir deploy/infrastructure/postgres/migrations postgres 'host=localhost port=5442 user=user password=passwd database=test sslmode=disable' up
pgstop:
	cd deploy/infrastructure/postgres && docker compose down && rm -rf postgres

run:
	go run services/contact/cmd/app/main.go