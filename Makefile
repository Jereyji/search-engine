
.PHONY: generate_env
generate_env:
	./scripts/generate_env.sh

.PHONY: migrate_up
migrate_up:
	migrate -path ./migrations -database 'postgres://postgres:1234@localhost:5436/postgres?sslmode=disable' up

.PHONY: migrate_down
migrate_down:
	migrate -path ./migrations -database 'postgres://postgres:1234@localhost:5436/postgres?sslmode=disable' down