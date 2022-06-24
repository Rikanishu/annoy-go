test:
	go test -v ./...

lint:
	golangci-lint run ./...

run-example:
	go run ./cmd/example/

build:
	go build -v ./...

goimports:
	goimports -w ./

gofmt:
	gofmt -w ./

verify: goimports gofmt build test