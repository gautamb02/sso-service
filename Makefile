# Makefile for go-common

APP_NAME := sso-service
BIN_DIR := bin
BIN_FILE := $(BIN_DIR)/$(APP_NAME)
CONFIG ?= /home/gautam/Desktop/DevHub/PersonalHub/sso-service/resource/config/config.yml

.PHONY: all build run test lint clean help

all: build

## Build the application into bin/
build: $(BIN_DIR)
	@echo "🚀 Building..."
	go build -o $(BIN_FILE) .

## Ensure bin directory exists
$(BIN_DIR):
	mkdir -p $(BIN_DIR)

## Run the application with config
run: build
	@echo "🚀 Running with config '$(CONFIG)'..."
	$(BIN_FILE) --config=$(CONFIG)

## Run tests
test:
	@echo "🧪 Running tests..."
	go test ./...

## Lint the code
lint:
	@echo "🔍 Linting..."
	golangci-lint run || echo "Install golangci-lint to lint properly."

## Clean build artifacts
clean:
	@echo "🧹 Cleaning up..."
	rm -rf $(BIN_DIR)

## Show help
help:
	@echo ""
	@echo "🛠️  Available targets:"
	@echo "    make build             - Build the application into $(BIN_DIR)/"
	@echo "    make run CONFIG=...    - Run with config file (default: config.yml)"
	@echo "    make test              - Run tests"
	@echo "    make lint              - Run linter"
	@echo "    make clean             - Clean binaries"
	@echo "    make help              - Show this help"
	@echo ""
