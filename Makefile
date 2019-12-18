.PHONY: build run test test-cover lint

build:
	@mkdir -p bin
	@go build -o bin/alien-invasion \
        -ldflags "-X main.version=$${VERSION:-$$(git describe --tags --always --dirty)}" \
        ./cmd/main.go

test:
	@go test -timeout=1m ./...

test-cover:
	@rm -f all.coverage.out
	@go test -race -v -timeout=1m \
        -coverprofile=all.coverage.out \
        -coverpkg=./... $$(go list ./...|grep -v cmd)

lint:
	docker pull golangci/golangci-lint:latest
	docker run -v`pwd`:/workspace -w /workspace \
        golangci/golangci-lint:latest golangci-lint run ./...
