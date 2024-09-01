# Define the variables
PROJECT_NAME = url-shortener

# Commands to run
.PHONY: help tidy run test docker-build

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

docker-build:
	@echo "Building Docker Image..."
	docker build -t url-shortener .

docker-run:
	@echo "Building Docker Container..."
	docker run --rm -it --env-file .env -p $(PORT):8080 --name url-shortenerapp url-shortener

include .env
export