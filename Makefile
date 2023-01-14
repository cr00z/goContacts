pgstart:
	cd deploy/infrastructure/postgres && docker compose up -d
	sleep 3
	docker ps


pgstop:
	cd deploy/infrastructure/postgres && docker compose down && rm -rf postgres

run:
	go run services/contact/cmd/app/main.go