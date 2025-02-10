.PHONY: clean critic security lint test build generate run

APP_NAME = apiserver
CUR_DIR = $(PWD)
BUILD_DIR = $(PWD)/build

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
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install github.com/rubenv/sql-migrate/...@latest
	cp $(GOPATH)/bin/sql-migrate $(BUILD_DIR)/sql-migrate
	mkdir -p $(BUILD_DIR)/platform/migrations
	cp -R $(PWD)/platform/migrations $(BUILD_DIR)/platform
	cp -R $(PWD)/migrate.sh $(BUILD_DIR)
	cp -R $(PWD)/start.sh $(BUILD_DIR)

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

crud:
	cd ../gen && go run . -model $(MODEL) -root "gitlab.com/a10869/api-modules/backend" -directory $(CUR_DIR) && cd $(CUR_DIR)
