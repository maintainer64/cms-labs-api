.PHONY: clean critic security lint test generate build run

APP_NAME = gen
BUILD_DIR = $(PWD)/build

clean:
	rm -rf ./build

critic:
	gocritic check -enableAll ./...

security:
	gosec ./...

lint:
	golangci-lint run ./... --timeout=10m

generate:
	go generate

test:
	go test -v -timeout 30s -coverprofile=cover.out -cover ./...
	go tool cover -func=cover.out
