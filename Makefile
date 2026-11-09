.PHONY: build run 

build:
	@go build -o bin/core   ./cmd/core

run:build
	@./bin/core