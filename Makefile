
.PHONY: create_env
env:
	./scripts/create_env.sh

.PHONY: docker_run
docker_run:
	docker run --name=todo-db -e POSTGRES_PASSWORD='1234' -p 5436:5432 -d postgres

.PHONY: migrate_up
migrate_up:
	migrate -path ./migrations -database 'postgres://postgres:1234@localhost:5436/postgres?sslmode=disable' up

.PHONY: migrate_down
migrate_down:
	migrate -path ./migrations -database 'postgres://postgres:1234@localhost:5436/postgres?sslmode=disable' down



PROJECT_NAME=search-engine

MIGRATE_CMD := $(shell which migrate)
MIGRATE_VERSION := v4.15.2

SERVICE_PATH=$(CURDIR)/cmd/main.go
MIGRATION_PATH=$(CURDIR)/migrations

SCRIPTS_PATH=$(CURDIR)/scripts
EXPORT_DEFAULT_ENV_SCRIPT=$(SCRIPTS_PATH)/export_default_env.sh

DEPLOYMENTS_PATH=$(CURDIR)/deployments
DOCKER_COMPOSE_PATH=$(DEPLOYMENTS_PATH)/docker-compose.yaml

DATABASE_URL=postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_DB)?sslmode=$(DB_SSLMODE)

MIGRATION_NAME := $(name)


build-service:
	docker-compose -p $(PROJECT_NAME) -f $(DOCKER_COMPOSE_PATH) build 

up-service: 
	docker-compose -p $(PROJECT_NAME) -f $(DOCKER_COMPOSE_PATH) up -d

down-service:
	docker-compose -p $(PROJECT_NAME) -f $(DOCKER_COMPOSE_PATH) down

run-service:
	go run $(SERVICE_PATH)

.check-migrate:
	ifeq ($(MIGRATE_CMD),)
		@echo "migrate not found, installing..."
		go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	endif

migration-up: .check-migrate
	@echo "Running migrations up..."
	migrate -path $(MIGRATION_PATH) -database $(DATABASE_URL) up

migration-down: .check-migrate
	@echo "Running migrations down..."
	migrate -path $(MIGRATION_PATH) -database $(DATABASE_URL) down

migration-create: .check-migrate
	@if [ -z "$(MIGRATION_NAME)" ]; then \
		echo "Migration name is required. Use: make migration-create name=<migration_name>"; \
		exit 1; \
	fi
	@echo "Creating new migration: $(MIGRATION_NAME)"
	migrate create -ext sql -dir $(MIGRATION_PATH) -seq $(MIGRATION_NAME)

.PHONY: lines-count
lines-count:
	@echo  Number of lines in GO files:
	@echo  ""[${shell find $(CURDIR) -name '*.go' -type f -print0 | xargs -0 cat | wc -	}]