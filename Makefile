install-codegen:
	go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest

generate-api:
	go generate ./...

run:
	go build -o main .
	./main