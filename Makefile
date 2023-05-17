BINARY_NAME = main

build:
	go build -o bin/$(BINARY_NAME) -v ./cmd/main.go

run:
	go run cmd/main.go