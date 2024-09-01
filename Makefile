# Define the variables
PROJECT_NAME = url-shortener

# Commands to run
.PHONY: help tidy run test

help:
	@echo "Usage:"
	@echo "  make <target>"
	@echo ""
	@echo "Targets:"
	@echo "  tidy      Format the code, vendor, and tidy the Go modules"
	@echo "  run       Run the Go application locally without Docker."
	@echo "  test      Runs all the tests in the sub folders."

tidy:
	go fmt ./...
	go mod tidy
	go mod vendor

run:
	go run cmd/server/main.go

test:
	go test ./...

include .env
export