include .env
MIGRATE=docker-compose exec web migrate -path=migration -database "mysql://${DB_USERNAME}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}" -verbose

migrate-up:
		$(MIGRATE) up
migrate-down:
		$(MIGRATE) down

drop:
		$(MIGRATE) drop

create:
		@read -p  "What is the name of migration?" NAME; \
		${MIGRATE} create -ext sql -seq -dir migration  $$NAME

### Others are not necessary