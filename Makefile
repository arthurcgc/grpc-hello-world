.PHONY: protos server

server:
	go build cmd/main.go
	./main

build-grpc:
	protoc --go_out=. --go_opt=paths=source_relative protos/hello/hello.proto
	protoc --go-grpc_out=. --go-grpc_opt=paths=source_relative protos/hello/hello.proto

