
migrate-up:
	migrate -path ./migrations/database -database "postgres://sykros:fqQ3nN4L@localhost:9001/sykros_files?sslmode=disable" up