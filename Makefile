run:
	go run cmd/server/main.go

migrations:
	go run db/migrate/migrate.go

swag:
	cd cmd/server; swag init --parseDependency --parseInternal --parseDepth 1
