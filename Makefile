dev: gen
	go run .
test:
	go test ./...
gen:
	git rm -f --ignore-unmatch **/*fakes/*fake*.go
	go generate ./...
	git add **/*fakes/*fake*.go
lint:
	find . -type f -name "*.templ" -exec templ fmt "{}" \;
	gofumpt -d -w .
	golangci-lint run
	swag fmt
validate: lint test

done: gen validate

.PHONY: run test gen lint validate done
