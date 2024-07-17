db-migrate-up:
	migrate -path ./db/migration -database "$(POSTGRES_URL)" -verbose up

db-migrate-down:
	migrate -path ./db/migration -database "$(POSTGRES_URL)" -verbose down $(STEPS)

run:
	go run ./cmd/server"postgres://${}:postgres@localhost:5432/postgres?sslmode=disable"