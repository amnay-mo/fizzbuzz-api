.DEFAULT_GOAL=noop
.DELETE_ON_ERROR:

.PHONY: noop
noop:


GOCMD        := go
GOBUILD      := $(GOCMD) build
GOCLEAN      := $(GOCMD) clean
GOTEST       := $(GOCMD) test
GOGET        := $(GOCMD) get
GOMETALINTER := GOGC=400 gometalinter --disable-all -E deadcode -E errcheck -E gocyclo -E gofmt -E goimports -E golint -E ineffassign -E megacheck -E misspell -E nakedret -E structcheck -E unconvert -E unparam -E varcheck -E vet\
 --tests --vendor --warn-unmatched-nolint --sort=path --sort=line --deadline=10m --concurrency=2 --enable-gc ./...


.PHONY: install-gometalinter
install-gometalinter:
	$(GOGET) github.com/alecthomas/gometalinter
	gometalinter --install > /dev/null

.PHONY: lint
lint: install-gometalinter
	$(GOMETALINTER)

.PHONY: test
test:
	$(GOTEST) -v ./fizzbuzz ./api

.PHONY: build
build:
	$(GOBUILD) -o build/fizzbuzz-api

clean::
	rm -rf build
