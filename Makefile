# Defaults: macOS (darwin/arm64)
APP_NAME ?= api
GOOS ?= darwin
GOARCH ?= arm64

.PHONY: help ## Prints help for targets with comments
.DEFAULT_GOAL := help
help:
	@echo "Usage: make [target]"
	@echo "Targets:"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z0-9_-]+:.*?## / {printf "  %-20s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: dev
dev: ## Run the application in development mode
	@air -c .air.toml

.PHONY: build
build: ## Build the application
	CGO_ENABLED=1 GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o bin/$(APP_NAME)-$(GOOS)-$(GOARCH) cmd/$(APP_NAME)/main.go

.PHONY: build.api
build.api: ## Build the api application
	@@make build APP_NAME=api

.PHONY: build.worker
build.worker: ## Build the worker application
	@@make build APP_NAME=worker

.PHONY: build.linux
build.linux: ## Build the application for linux
	@@make build GOOS=linux GOARCH=amd64

.PHONY: lint
lint: ## Run linter
	@golangci-lint run

.PHONY: organize-imports
organize-imports: ## Organize imports
	@command -v goimports-reviser > /dev/null || go install -v github.com/incu6us/goimports-reviser/v3@latest && exit 0
	@goimports-reviser -rm-unused -set-alias -format -recursive ./...

.PHONY: fmt
fmt: ## Format the code
	@gofmt -w -s ./internal ./cmd 

ci: lint organize-imports fmt ## Run all checks

.PHONY: install-tools
install-tools: ## Install tools
	@command -v air > /dev/null || go install -v github.com/cosmtrek/air@latest && exit 0
	@command -v golangci-lint > /dev/null || go install -v github.com/golangci/golangci-lint/cmd/golangci-lint@latest && exit 0
	@command -v goimports-reviser > /dev/null || go install -v github.com/incu6us/goimports-reviser/v3@latest && exit 0
	@command -v gomock > /dev/null || go install -v github.com/golang/mock/mockgen@latest && exit 0
	@command -v lefthook > /dev/null || go install -v github.com/evilmartians/lefthook@latest && exit 0
