.PHONY: build test integration lint

build:
	CGO_ENABLED=0 go build -trimpath -o bin/ ./cmd/...

test:
	go test ./internal/...

integration:
	go test -v -count=1 ./integration/...

lint:
	golangci-lint run
