include local.env
export

run:
	@echo "Running API..."
	go run ./cmd/api
migration:
	@echo "Creating migration files for ${name}..."
	migrate create -ext=.sql -dir=./migrations ${name}
migrate_up:
	    migrate -path=./migrations -database=postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE) up
migrate_down:
	    migrate -path=./migrations -database=postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE) down
test:
	go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out