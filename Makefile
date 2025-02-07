.PHONY: clean critic security lint test generate build run

BUILD_DIR = $(PWD)/build
GO_PACKAGES = ./backend/... ./gen/... ./shared/... ./pnetlabaddon/...

clean:
	rm -rf ./build

critic:
	gocritic check -enableAll ./...

security:
	gosec ./...

lint:
	golangci-lint run ${GO_PACKAGES}

pre_commit:
	pre-commit run --all-files

generate:
	make -C backend generate
	make -C pnetlabaddon generate
	yarn --cwd nextui-dashboard generate

test: clean critic security lint pre_commit
	go test -v -coverprofile=coverage.out -covermode=count $(GO_PACKAGES) > tests.out 2>&1; \
	TEST_EXIT_CODE=$$?; \
	cat tests.out; \
	if [ $$TEST_EXIT_CODE -ne 0 ]; then \
		echo "testing failed" >&2; \
		exit $$TEST_EXIT_CODE; \
	fi

dev:
	./CI-CD/dev.sh
