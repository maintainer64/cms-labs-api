.PHONY: critic security lint test generate

APP_NAME = shared
CUR_DIR = $(CURDIR)
BUILD_DIR = $(CURDIR)/build

critic:
	gocritic check -enableAll ./...

security:
	gosec ./...

lint:
	golangci-lint run ./... --timeout=10m

test:
	go test -v -timeout 30s -coverprofile=cover.out -cover ./...
	go tool cover -func=cover.out

generate:
	go generate
