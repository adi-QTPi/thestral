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