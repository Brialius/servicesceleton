VERSION ?= $(shell git describe --tags --dirty --always --match=v* || echo v0)
BUILD := $(shell git rev-parse --short HEAD)
LDFLAGS=-ldflags "-X=main.version=$(VERSION) -X=main.build=$(BUILD)"
BUILDFLAGS=$(LDFLAGS)
PROJECTNAME=servicesceleton
GOEXE := $(shell go env GOEXE)
BIN=bin/$(PROJECTNAME)$(GOEXE)

.PHONY: setup
setup: ## Install all the build and lint dependencies
#	go get -m github.com/golangci/golangci-lint/cmd/golangci-lint

.PHONY: test
test: ## Run all the tests
	go test -v $(BUILDFLAGS) ./...

.PHONY: lint
lint: ## Run all the linters
	golangci-lint run --enable-all --disable gochecknoinits --disable gochecknoglobals --disable goimports \
	--out-format=tab --tests=false ./...

.PHONY: ci
ci: setup lint test build ## Run all the tests and code checks

.PHONY: build
build: mod-refresh ## Build a version
	go build $(BUILDFLAGS) -o $(BIN)

.PHONY: install
install: mod-refresh ## Install a binary
	go install $(BUILDFLAGS)

.PHONY: clean
clean: ## Remove temporary files
	go clean

.PHONY: mod-refresh
mod-refresh: clean ## Refresh modules
	go mod tidy -v

.PHONY: version
version:
	@echo $(VERSION)-$(BUILD)

.PHONY: release
release:
	git tag $(ver)
	git push origin --tags

.DEFAULT_GOAL := build
