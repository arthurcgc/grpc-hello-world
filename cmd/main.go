package main

import (
	"log"
	"net"

	"github.com/arthurcgc/grpc-learning/protos/currency"
	"github.com/arthurcgc/grpc-learning/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	gs := grpc.NewServer()
	cs := server.New()
	currency.RegisterCurrencyServer(gs, cs)
	reflection.Register(gs)

	listener, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(gs.Serve(listener))
}
