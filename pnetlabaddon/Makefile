.PHONY: clean critic security lint test generate build run

APP_NAME = pnetlabaddon
CUR_DIR = $(PWD)
BUILD_DIR = $(PWD)/pnetlabaddon/usr/local/bin

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
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install github.com/rubenv/sql-migrate/...@latest
	cp ${HOME}/go/bin/sql-migrate $(PWD)/pnetlabaddon/usr/local/bin/
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -a -installsuffix cgo -o $(BUILD_DIR)/$(APP_NAME) main.go
	cp -R $(PWD)/platform/migrations $(PWD)/pnetlabaddon/etc/pnetlabaddon
	cp -R $(PWD)/migrate.sh $(PWD)/pnetlabaddon/etc/pnetlabaddon

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
	cd ../gen && go run . -model $(MODEL) -root "gitlab.com/a10869/api-modules/pnetlabaddon" -directory $(CUR_DIR) && cd $(CUR_DIR)
