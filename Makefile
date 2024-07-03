install-deps:
	go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
	go install go.uber.org/mock/mockgen@latest

generate:
	go generate ./...

test:
	go test ./...

build: generate
	go build -o main .

run:
	./main