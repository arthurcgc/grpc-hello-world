package main

import (
	"log"
	"net"

	"github.com/arthurcgc/grpc-learning/protos/hello"
	"github.com/arthurcgc/grpc-learning/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	gs := grpc.NewServer()
	helloServer := server.New()
	hello.RegisterHelloServer(gs, helloServer)
	reflection.Register(gs)

	listener, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(gs.Serve(listener))
}
