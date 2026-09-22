.PHONY: clean critic security lint test build generate run

APP_NAME = clabgate
CUR_DIR = $(CURDIR)
BUILD_DIR = $(CURDIR)/build

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
	cp -R "$(CURDIR)/start.sh" "$(BUILD_DIR)"

run: generate build
	$(BUILD_DIR)/$(APP_NAME)

generate:
	go generate
