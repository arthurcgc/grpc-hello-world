.PHONY: protos

build-grpc:
	protoc --go_out=. --go_opt=paths=source_relative protos/currency.proto
	protoc --go-grpc_out=. --go-grpc_opt=paths=source_relative protos/currency.proto