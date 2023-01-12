pgstart:
	cd deploy/infrastructure/postgres && docker compose up -d

pgstop:
	cd deploy/infrastructure/postgres && docker compose down && rm -rf postgres

run:
	go run services/contact/cmd/app/main.go