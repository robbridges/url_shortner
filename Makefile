include local.env
export

migration:
	@echo "Creating migration fo;es fpr ${name}..."
	migrate create -ext=.sql -dir=./migrations ${name}
migrate_up:
	    migrate -path=./migrations -database=postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE) up
migrate_down:
	    migrate -path=./migrations -database=postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE) down