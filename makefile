APP_NAME?=thestral
MAIN_FILE?=./cmd/main.go
PORT?=80
ADMIN_PORT?=7008
SECURE_IP?=100.113.160.66
PLATFORM?=linux/amd64

.PHONY: start run build clean test docker-build docker-run

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
	@docker run -d --network host -e PORT=$(PORT) -e ADMIN_PORT=$(ADMIN_PORT) -e SECURE_IP=$(SECURE_IP) --rm --name $(APP_NAME)-container $(APP_NAME)