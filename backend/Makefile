.PHONY: clean critic security lint test build generate run

APP_NAME = apiserver
CUR_DIR = $(CURDIR)
BUILD_DIR = $(CURDIR)/build
SQL_MIGRATE_VERSION = v1.8.1

clean:
	rm -rf ./build

critic:
	gocritic check -enableAll ./...

security:
	gosec ./...

lint:
	golangci-lint run ./... --timeout=10m

test:
	go test -v -timeout 30s -coverprofile=cover.out -cover ./...
	go tool cover -func=cover.out

build: clean
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -a -installsuffix cgo -o "$(BUILD_DIR)/$(APP_NAME)" main.go
	GOBIN="$(BUILD_DIR)" CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install github.com/rubenv/sql-migrate/...@$(SQL_MIGRATE_VERSION)
	mkdir -p $(BUILD_DIR)/platform/migrations
	cp -R $(CURDIR)/platform/migrations $(BUILD_DIR)/platform
	cp -R $(CURDIR)/migrate.sh $(BUILD_DIR)
	cp -R $(CURDIR)/start.sh $(BUILD_DIR)

run: generate build
	$(BUILD_DIR)/$(APP_NAME)

migrate.up:
	./migrate.sh up

migrate.down:
	./migrate.sh d

migrate.status:
	./migrate.sh s

generate:
	go generate
