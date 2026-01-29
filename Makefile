.PHONY: build test integration lint

PLATFORMS = linux darwin
ARCHITECTURES = amd64 arm64

build:
	CGO_ENABLED=0 go build -trimpath -o bin/ ./cmd/...

build-all:
	@for os in $(PLATFORMS); do \
		for arch in $(ARCHITECTURES); do \
			echo "Building for $$os/$$arch..."; \
			CGO_ENABLED=0 GOOS=$$os GOARCH=$$arch go build -trimpath -o bin/$$os-$$arch/ ./cmd/...; \
		done; \
	done

test:
	go test ./internal/...

integration:
	go test -v -count=1 ./integration/...

lint:
	golangci-lint run
