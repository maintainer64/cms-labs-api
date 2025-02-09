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
	make -C gen generate
	make -C pnetlabaddon generate
	make -C shared generate
	yarn --cwd nextui-dashboard generate

test: clean critic security lint pre_commit
	go test -v -coverprofile=coverage.out -covermode=count $(GO_PACKAGES) > tests.out 2>&1; \
	TEST_EXIT_CODE=$$?; \
	cat tests.out; \
	if [ $$TEST_EXIT_CODE -ne 0 ]; then \
		echo "testing failed" >&2; \
		exit $$TEST_EXIT_CODE; \
	fi

generate_certs:
	@sudo openssl req -x509 -nodes -days 3650 -newkey rsa:4096 -keyout ./certs/private.key -out ./certs/public.crt
	@sudo chmod -R 755 ./certs


dev:
	./CI-CD/dev.sh
