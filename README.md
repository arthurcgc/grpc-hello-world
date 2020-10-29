# A hello world example for a grpc server

## Instructions
* run ```make build-grpc``` to generate the grpc and interface codes
* implement the HelloServer interface, specified at protos/hello/hello_grpc.go
* create a tcp listener, instantiate a grpc server with grpc.NewServer(), instantiate the implementation of HelloServer, bind the grpc server and the HelloServer implementation with: hello.RegisterHelloServer()
* call grpc Serve method :)

## Running
* ```make server```