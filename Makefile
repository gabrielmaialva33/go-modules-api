.PHONY: test lint build run

BINARY_NAME=go-modules-api
MAIN=main.go
BIN_DIR=bin

test:
	go test -v ./...

lint:
	golangci-lint run

build:
	@mkdir -p $(BIN_DIR)
	go mod tidy
	go build -o $(BIN_DIR)/$(BINARY_NAME) $(MAIN)

run: build
	./$(BIN_DIR)/$(BINARY_NAME)

clean:
	go clean
	rm -f $(BIN_DIR)/$(BINARY_NAME)