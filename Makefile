.PHONY: clean critic security lint test build run

APP_NAME = gen
BUILD_DIR = $(PWD)/build
GO_PACKAGES = ./backend/... ./gen/...

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
	go test -v -coverprofile=coverage.out -covermode=count $(GO_PACKAGES) > tests.out 2>&1; \
	TEST_EXIT_CODE=$$?; \
	cat tests.out; \
	if [ $$TEST_EXIT_CODE -ne 0 ]; then \
		echo "testing failed" >&2; \
		exit $$TEST_EXIT_CODE; \
	fi
