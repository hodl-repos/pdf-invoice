.PHONY: build

run:
	@go run ./cmd/qr

build:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o=./bin/ ./cmd/...

test:
	@go test \
		-count=1 \
		-cover \
		-timeout=10s \
		./...