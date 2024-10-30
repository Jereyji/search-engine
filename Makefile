PROJECT_NAME=search-engine
SERVICE_PATH=$(CURDIR)/cmd/main.go
DEPLOYMENTS_PATH=$(CURDIR)/deployments
DOCKER_COMPOSE_PATH=$(DEPLOYMENTS_PATH)/docker-compose.yaml

.PHONY: env
env:
	./scripts/create_env.sh

.PHONY: up-service
up-service: 
	docker-compose -p $(PROJECT_NAME) -f $(DOCKER_COMPOSE_PATH) up -d

.PHONY: down-service
down-service:
	docker-compose -p $(PROJECT_NAME) -f $(DOCKER_COMPOSE_PATH) down

.PHONY: run-service
run-service:
	go run $(SERVICE_PATH)
