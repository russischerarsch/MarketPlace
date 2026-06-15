include .env
export

migrate-up:
	migrate -path migrations -database ${CONN_STRING} up 1

migrate-down:
	migrate -path migrations -database ${CONN_STRING} down 1