all: generate_grpc_code clean build

generate_grpc_code:
	protoc --go_out=plugins=grpc:. api/api.proto

clean:
	go clean

build:
	go build -o bin/server

run:
	bin/server

# Testing
.PHONY: run test integration
test:
	go test -v -count=1 ./...

integration:
	go test -v -count=1 -tags integration ./...
