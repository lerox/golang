.PHONY: test build format

test:
	go test ./...

format:
	@bash ./scripts/fmt-check.sh

build:
	go build -o ./bin/main -race -v ./cmd/main.go
