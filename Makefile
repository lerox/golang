.PHONY: test build

test:
	go test ./...

build:
	go build -o ./bin/lerox-golang -race -v ./cmd/main.go
