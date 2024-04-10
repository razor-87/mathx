lint:
	@golangci-lint run

test:
	@go test ./...

test-shfl:
	@go test -shuffle=on -timeout=5s -count 3 ./...

test-race:
	@go test -race -timeout=10s -count 1 ./...

.PHONY: lint test test-slow test-race
