.DEFAULT_GOAL=noop
.DELETE_ON_ERROR:

.PHONY: noop
noop:

GO111MODULE=on


GOCMD        := go
GOBUILD      := $(GOCMD) build
GOCLEAN      := $(GOCMD) clean
GOTEST       := $(GOCMD) test
GOGET        := $(GOCMD) get


.PHONY: lint
lint: golangci-lint

GOLANGCI_LINT_VERSION=v1.17.1
GOLANGCI_LINT_DIR=$(shell go env GOPATH)/pkg/golangci-lint/$(GOLANGCI_LINT_VERSION)
$(GOLANGCI_LINT_DIR):
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(GOLANGCI_LINT_DIR) $(GOLANGCI_LINT_VERSION)

.PHONY: install-golangci-lint
install-golangci-lint: $(GOLANGCI_LINT_DIR)

.PHONY: golangci-lint
golangci-lint: install-golangci-lint
	$(GOLANGCI_LINT_DIR)/golangci-lint run --disable-all \
		--exclude-use-default=false \
		--enable=govet \
		--enable=ineffassign \
		--enable=deadcode \
		--enable=golint \
		--enable=goconst \
		--enable=gofmt \
		--enable=goimports \
		--skip-dirs=pkg/client/ \
		--deadline=120s \
		--tests ./...


.PHONY: test
test:
	$(GOTEST) -v ./...

.PHONY: build
build:
	$(GOBUILD) -o build/fizzbuzz-api

.PHONY: get
get:
	$(GOGET)

clean::
	rm -rf build
