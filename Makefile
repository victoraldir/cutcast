GOCMD:=$(shell which go)
GOFMT:=$(shell which gofmt)
GOBUILD:=$(GOCMD) build
GODOWNLOAD:=$(GOCMD) mod download
GOINSTALL:=$(GOCMD) install
GOCLEAN:=$(GOCMD) clean
GOTEST:=$(GOCMD) test
GOGET:=$(GOCMD) get
GOLIST:=$(GOCMD) list
GOVET:=$(GOCMD) vet
GOPATH:=$(shell $(GOCMD) env GOPATH)

all: test build

.PHONY: check-deps
check-deps:
	@echo "docker: $$(which docker || echo 'not found')"
	@echo "go: $$(which go || echo 'not found')"
	@echo "GolangCI-Lint: $$(which golangci-lint || echo 'not found')"

.PHONY: build
build: deps
	$(GOBUILD) -o $(BINARY_NAME) ./cutcast

.PHONY: install
install: deps
	$(GOINSTALL) ./

.PHONY: deps
deps:
	$(GODOWNLOAD)

.PHONY: lint
lint:
	@golangci-lint run -v --timeout=5m -E gosec -E revive -E goconst -E misspell -E whitespace ./...