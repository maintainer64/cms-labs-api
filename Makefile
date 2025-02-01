.PHONY: clean critic security lint test build run

APP_NAME = gen
BUILD_DIR = $(PWD)/build


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
	make -C backend test
	make -C gen test
