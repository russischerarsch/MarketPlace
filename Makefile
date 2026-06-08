include .env
export

migrate-up:
	migrate -path migrations -database "postgres://postgres:dkfl26052010@localhost:5432/postgres?sslmode=disable" up 1

migrate-down:
	migrate -path migrations -database ${CONN_STRING} down 1