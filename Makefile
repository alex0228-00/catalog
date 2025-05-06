# build go binary
BUILD_DIR := bin
APP_NAME := catalog
SERVER_ENTRY := src/cmd/server/main.go

.PHONY: all build clean run ut ut-db docker-up docker-down docker-clean

all: build

build:
	@echo "Building $(APP_NAME)..."
	@mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=amd64 go build -tags=mysql -o $(BUILD_DIR)/$(APP_NAME) $(SERVER_ENTRY)

clean:
	@echo "Cleaning build directory..."
	@rm -rf $(BUILD_DIR)

# Unit tests
UT_DB_TAGS ?= database,mysql

ut:
	go test -v $(go list ./src/... | grep -v './src/cmd' | grep -v './src/datastore/ent')

ut-db:
	go test -v -tags=$(UT_DB_TAGS) ./src/service/...

# Docker compose for tests
COMPOSE_FILE=scripts/docker-compose-test.yml

docker-up:
	docker compose -f $(COMPOSE_FILE) up -d

docker-down:
	docker compose -f $(COMPOSE_FILE) down

docker-clean:
	docker compose -f $(COMPOSE_FILE) down -v