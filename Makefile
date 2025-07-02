DB_HOST ?= localhost
DB_USER ?= faisal
DB_PASSWORD ?= faisalg5
DB_PORT ?= 5432
DB_NAME ?= cinevodb
DB_SSLMODE ?= disable

MIGRATION_DIR ?= migrations

MIGRATE = migrate -source file://$(MIGRATION_DIR) -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)"

migration_create:
	migrate create -seq -dir $(MIGRATION_DIR) -ext sql $(name)

migration_up:
	$(MIGRATE) up

migration_down:
	$(MIGRATE) down 1

migration_drop:
	$(MIGRATE) drop -f

migration_force:
	$(MIGRATE) force $(ver)