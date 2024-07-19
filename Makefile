docker-up:
	docker compose --env-file ./.env up -d

db-migrate-up:
	migrate -path ./db/migration -database "$(POSTGRES_URL)" -verbose up

db-migrate-down:
	migrate -path ./db/migration -database "$(POSTGRES_URL)" -verbose down $(STEPS)

sqlc:
	sqlc generate

run:
	go run ./cmd/server

test:
	go test -v -cover ./...