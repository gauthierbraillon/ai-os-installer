.PHONY: build test lint run clean

build:
	go build -o bin/ai-os-installer ./cmd/ai-os-installer

test:
	go test -v ./...

lint:
	golangci-lint run

run:
	go run ./cmd/ai-os-installer

clean:
	rm -rf bin/
