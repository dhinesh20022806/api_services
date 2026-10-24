.PHONY: lint fmt test check

lint:
	golangci-lint run ./...

fmt:
	golangci-lint fmt

run:
	godotenv -f .env go run main.go

test:
	go test ./...

check: fmt lint test