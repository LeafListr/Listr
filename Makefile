dev:
	go run .
test:
	go test ./...
gen:
	swag init -g internal/api/api.go
	go generate ./...
lint:
	gofumpt -d -w .
	golangci-lint run
	swag fmt
validate: gen lint test

done: gen validate

.PHONY: run test gen lint validate done
