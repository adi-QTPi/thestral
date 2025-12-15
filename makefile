APP_NAME=thestral
MAIN_FILE=./cmd/main.go
PORT=7007
ADMIN_PORT=7008

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
	@echo "\n---Docker Build---\n"
	@docker buildx build -f docker/Dockerfile --build-arg APP_NAME=$(APP_NAME) -t $(APP_NAME) .
docker-run:
	@echo "\n---Docker Run Image---\n"
	@docker run \
		-p $(PORT):7007 --rm\
		-p $(ADMIN_PORT):7008 --rm\
		--name $(APP_NAME)-container\
		$(APP_NAME)