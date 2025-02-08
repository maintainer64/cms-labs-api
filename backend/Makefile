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
	golangci-lint run ./...

test:
	go test -v -timeout 30s -coverprofile=cover.out -cover ./...
	go tool cover -func=cover.out

build:
	CGO_ENABLED=0 go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) main.go

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
