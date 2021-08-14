GOLANGCI_VERSION := 1.33.0

BIN := $(abspath ./bin)
PACKAGES := $(shell go list ./... | grep -v -e "mocks" -v -e "proto")
TEST_PACKAGES := $(shell go list ./internal/... ./pkg/... | grep -v -e "pkg/metadata")

.PHONY: setup run build start

setup:
	wget -O - -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v${GOLANGCI_VERSION}
	go install github.com/golang/mock/gomock@latest
	go install github.com/golang/mock/mockgen@latest

run:
	go run ./cmd/main.go

build:
	go build -o app ./cmd/main.go

start:
	./app

fmt:
	! gofmt -d -s . | grep '^'

vet:
	go vet ./... $(PACKAGES)

lint:
	$(BIN)/golangci-lint run -c .golangci.yaml ./...

test:
	go test -v -cover $(TEST_PACKAGES)

mockgen:
	$(BIN)/build-mocks
