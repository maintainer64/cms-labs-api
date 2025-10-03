.PHONY: clean security pre_commit test generate dev

GO_PACKAGES = ./backend/... ./shared/... ./pnetlabaddon/... ./clabgate/...

clean:
	rm -rf ./build

security: clean
	gocritic check -enableAll ./...
	gosec ./...
	golangci-lint run ${GO_PACKAGES} --timeout=10m

pre_commit: clean
	pre-commit run --all-files

generate:
	make -C backend generate
	make -C gen generate
	make -C clabgate generate
	make -C pnetlabaddon generate
	make -C shared generate
	yarn --cwd nextui-dashboard generate

test: clean
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
