help:
	@printf "Usage: make [target] [VARIABLE=value]\nTargets:\n"
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

install: ## Install pre-commit hooks
	@pre-commit install
	@pre-commit gc

uninstall: ## Uninstall hooks
	@pre-commit uninstall

validate: ## Validate files with pre-commit hooks
	@pre-commit run --all-files

.PHONY: dependency
dependency: ## Dependency maintanance
	@go get -u ./...
	@go mod tidy

.PHONY: fmt
fmt: ## gofmt (reformat) package sources
	@go fmt $(shell go list ./...)

.PHONY: test
test: ## Run unit tests
	@go test ./... -v -cover

.PHONY: build
build: ## Compile packages and dependencies
	@go build ./cmd/gitlab-membership

.PHONY: build
run: ## Run locally
	@go run ./cmd/gitlab-membership
