package server

import (
	"context"

	"github.com/arthurcgc/grpc-learning/protos/hello"
	"github.com/sirupsen/logrus"
)

type internalServer struct {
	hello.UnimplementedHelloServer
	logger *logrus.Logger
}

func New() *internalServer {
	return &internalServer{logger: logrus.New()}
}

func (is *internalServer) SayHello(ctx context.Context, emp *hello.EmptyParam) (*hello.HelloResponse, error) {
	return &hello.HelloResponse{Hello: "Hello World!"}, nil
}
