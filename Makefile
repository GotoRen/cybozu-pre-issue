GOCMD=go
GORUN=$(GOCMD) run
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GODOC=$(GOCMD)doc
GOLANGCI=golangci-lint


all: env/apply run

.PHONY: build
build: ## go build
	$(GOBUILD) ./cmd/main.go


.PHONY: run
run: env/apply ## go run
	$(GORUN) ./cmd/main.go

.PHONY: env/apply
env/apply: ## apply environment variables
	@direnv allow

.PHONY: lint
lint: ## golang-ci lint
	$(GOLANGCI) run --config=.golangci.yaml ./...

.PHONY: doc
doc: ## godoc http:6060
	$(GODOC) -http=:6060



# Makefile config
#===============================================================
help:
	echo "Usage: make [task]\n\nTasks:"
	perl -nle 'printf("    \033[33m%-30s\033[0m %s\n",$$1,$$2) if /^([a-zA-Z0-9_\/-]*?):(?:.+?## )?(.*?)$$/' $(MAKEFILE_LIST)

.SILENT: help

.PHONY: $(shell egrep -o '^(\._)?[a-z_-]+:' $(MAKEFILE_LIST) | sed 's/://')
