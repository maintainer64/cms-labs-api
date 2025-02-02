.PHONY: clean critic security lint test build run

APP_NAME = gen
BUILD_DIR = $(PWD)/build
GO_PACKAGES = ./backend/... ./gen/...
GO_COVER_PACKAGES=$(go list $(GO_PACKAGES) | tr '\n' ',' | sed 's/,$//')

clean:
	rm -rf ./build

critic:
	gocritic check -enableAll ./...

security:
	gosec ./...

lint:
	make -C backend lint
	make -C gen lint

pre_commit:
	pre-commit run --all-files

test: clean critic security lint pre_commit
	go test -v -coverpkg="${GO_COVER_PACKAGES}" -coverprofile=coverage.out -covermode=count $(GO_PACKAGES) | tee `tests.txt`
	go tool cover -func coverage.out
