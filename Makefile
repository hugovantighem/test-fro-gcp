install-deps:
	go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
	go install go.uber.org/mock/mockgen@latest

generate:
	go generate ./...

generate-scratch: install-deps generate

test:
	go test ./...

test-scratch: generate-scratch test

build: 
	go build -o main .

build-scratch: generate-scratch build

run:
	./main

run-scratch: build-scratch run

copy-env-file:
	cp .env-example .env

start-db:
	docker run --name postgresql -e POSTGRES_USER=myusername -e POSTGRES_PASSWORD=mypassword -e POSTGRES_DB=mydb -p 5432:5432 -d postgres
