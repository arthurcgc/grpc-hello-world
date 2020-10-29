# A hello world example for a grpc server

## Instructions
* Run ```make build-grpc``` to generate the grpc and interface codes
* Implement the HelloServer interface, specified at protos/hello/hello_grpc.go
* Create a tcp listener, instantiate a grpc server with grpc.NewServer(), instantiate the implementation of HelloServer, bind the grpc server and the HelloServer implementation with: hello.RegisterHelloServer()
* Call grpc Serve method :)

## Running
* ```make server```

## Testing
* Once you got the server running you can do: ```grpcurl -plaintext localhost:9999 Hello.SayHello```