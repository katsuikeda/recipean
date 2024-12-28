BINARY_NAME=api
.DEFAULT_GOAL := run
.PHONY: build run clean test test-coverage test-coverage-out coverage vet security fmt lint vendor

all: test vet security fmt lint build

build:
	go build -mod=vendor -o ./bin/${BINARY_NAME} ./cmd/${BINARY_NAME}

run: build
	./bin/${BINARY_NAME}

clean:
	rm ./bin/${BINARY_NAME}

test:
	go test -mod=vendor ./...

test-coverage:
	go test -mod=vendor ./... -cover

test-coverage-report:
	mkdir -p ./coverage
	go test -mod=vendor ./... -coverprofile=./coverage/coverage.out

coverage:
	go tool cover -func=coverage/coverage.out -o coverage/coverage.html

vet:
	go vet ./...

security:
	gosec ./...

fmt:
	test -z $$(go fmt ./...)

lint:
	staticcheck ./...

vendor:
	go mod tidy
	go mod vendor