run:
	go run cmd/server/main.go

migrations:
	go run db/migrate/migrate.go