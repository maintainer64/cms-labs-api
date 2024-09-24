.PHONY: clean critic security lint test build run

APP_NAME = gen
BUILD_DIR = $(PWD)/build


clean:
	rm -rf ./build

critic:
	gocritic check -enableAll ./...

security:
	gosec ./...

lint:
	golangci-lint run ./...

pre_commit:
	pre-commit run --all-files

test: clean critic security lint pre_commit
	go test -v -timeout 30s -coverprofile=cover.out -cover ./...
	go tool cover -func=cover.out
