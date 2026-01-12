# app configurations
APP_NAME=thestral
MAIN_FILE=./cmd/main.go
PROXY_BIND?=0.0.0.0:80
ADMIN_BIND=100.113.160.66:7007
REDIS_PORT?=6381
REDIS_HOST?=azkaban
REDIS_PASSWORD?=super-secret-string
PLATFORM=linux/amd64

.PHONY: start run build clean test docker-build docker-run redis

start:
	@echo "\n---Running Live Reload---\n"
	@air
run:
	@echo "\n---Running---\n"
	@go run $(MAIN_FILE)
build:
	@echo "\n---Building---\n"
	@mkdir -p bin
	@go build -o bin/$(APP_NAME) $(MAIN_FILE)
clean:
	@echo "\n---Cleaning Binaries---\n"
	@rm -rf bin/
test:
	@echo "\n---Testing---\n"
	@go test ./...
docker-build:
	@echo "\n---Docker Build (Linux AMD64)---\n"
	@docker buildx build --platform $(PLATFORM) -f docker/Dockerfile --build-arg APP_NAME=$(APP_NAME) -t $(APP_NAME) .
docker-run:
	@echo "\n---Docker Run Image (Host Mode)---\n"
	@docker run -d --network host -e ADMIN_BIND=$(ADMIN_BIND) -e PROXY_BIND=$(PROXY_BIND) -e REDIS_PORT=$(REDIS_PORT) -e REDIS_HOST=$(REDIS_HOST) -e REDIS_PASSWORD=$(REDIS_PASSWORD) --rm --name $(APP_NAME)-container $(APP_NAME)
redis:
	@echo "\n---Spinning up new Redis Container---\n"
	@chmod +x ./redis/container.sh
	@./redis/container.sh

# db configurations
DB_CONTAINER=thestral_db
DB_NAME=thestral
DB_USER=admin
DB_PASSWORD=password
DB_PORT=5433

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