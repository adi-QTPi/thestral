# app configurations
APP_NAME:=thestral
MAIN_FILE:=./cmd/main.go
PROXY_BIND:=0.0.0.0:80
ADMIN_BIND:=100.113.160.66:7007

PLATFORM=linux/amd64

.PHONY: start run build clean test docker-build docker-run redis db-up db-down db-shell

start:
	@echo "\n---Running Live Reload---\n"
	@air
run:
	@echo "\n---Running---\n"
	@go run $(MAIN_FILE)
build:
	@echo "\n---Building---\n"
	@mkdir -p bin
	@go build -o bin/thestral ./cmd/main.go
clean:
	@echo "\n---Cleaning Binaries---\n"
	@rm -rf bin/
test:
	@echo "\n---Testing---\n"
	@go test ./...

V ?= dev
docker-build-push:
	@echo "\n---Docker Build (Linux AMD64)---\n"
	@docker buildx build --platform linux/amd64,linux/arm64 -f docker/Dockerfile -t adiqtpi/thestral:$(V) --push .
docker-run:
	@echo "\n---Docker Run Image (Host Mode)---\n"
	@docker run --network host -e ADMIN_BIND=$(ADMIN_BIND) -e PROXY_BIND=$(PROXY_BIND) -e DATABASE_URL=$(DATABASE_URL) -e DEBUG=$(DEBUG) --rm --name $(APP_NAME)-prod $(APP_NAME)

# db configurations
DB_HOST:=azkaban
DB_USER:=admin
DB_PROD_PASSWORD:=password
DB_NAME=thestral_db
DB_PROD_NAME:=thestral_prod
DB_PORT:=5433
DATABASE_URL := "host=$(DB_HOST) user=$(DB_USER) password=$(DB_PROD_PASSWORD) dbname=$(DB_PROD_NAME) port=$(DB_PORT) sslmode=disable TimeZone=UTC"

DB_CONTAINER=thestral_db

db-up:
	@if [ $$(docker ps -q -f name=$(DB_CONTAINER)) ]; then \
		echo "Database '$(DB_CONTAINER)' is already running."; \
	elif [ $$(docker ps -aq -f name=$(DB_CONTAINER)) ]; then \
		echo "Container exists but is stopped. Starting it..."; \
		docker start $(DB_CONTAINER); \
	else \
		echo "Creating and starting new Postgres container..."; \
		docker run -d \
			--name $(DB_CONTAINER) \
			-e POSTGRES_USER=$(DB_USER) \
			-e POSTGRES_PASSWORD=$(DB_PASSWORD) \
			-e POSTGRES_DB=$(DB_NAME) \
			-p $(DB_PORT):5432 \
			postgres:15-alpine; \
	fi
	
	@echo "Waiting for Postgres to accept connections..."
	@until docker exec $(DB_CONTAINER) pg_isready -U user; do \
		echo "Postgres is initializing..."; \
		sleep 1; \
	done

	@echo "Ensuring extensions are enabled..."
	@docker exec -i $(DB_CONTAINER) psql -U $(DB_USER) -d $(DB_NAME) -c 'CREATE EXTENSION IF NOT EXISTS "pgcrypto";' > /dev/null
	@echo "Database is ready!"

db-down:
	@echo "Stopping Postgres..."
	docker stop $(DB_CONTAINER) || true
	docker rm $(DB_CONTAINER) || true 

db-shell:
	docker exec -it $(DB_CONTAINER) psql -U $(DB_USER) -d $(DB_NAME)