.PHONY: build test integration lint run clean

build:
	go build -o bin/ai-os-installer ./cmd/ai-os-installer

test:
	go test -v -coverprofile=coverage.out ./...

integration:
	go test -v -tags=integration -timeout=5m ./tests/integration/...

lint:
	@which golangci-lint > /dev/null 2>&1 || (echo "golangci-lint not found: https://golangci-lint.run/usage/install/" && exit 1)
	golangci-lint run

run:
	go run ./cmd/ai-os-installer

clean:
	rm -rf bin/ coverage.out
