.PHONY: build

run:
	@go run ./cmd/service

req-sample-invoice:
	@curl -X POST -H "Content-Type: application/json" -d @./internal/service/handler/v1/sample.json localhost:12003/v1/generate > out/SampleInvoice.pdf

build:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o=./bin/ ./cmd/...

test:
	@go test \
		-count=1 \
		-cover \
		-timeout=10s \
		./...
