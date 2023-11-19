run:
	go run .
test:
	go test ./...
gen:
	go generate ./...
lint:
	gofumpt -d -w .
	golangci-lint run

.PHONY: gen test run