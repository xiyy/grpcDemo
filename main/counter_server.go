package main

import (
	"google.golang.org/grpc"
	"grpcDemo/counter"
	"net"
)

func main() {
	server := grpc.NewServer()
	counter.RegisterCounterServer(server, &counter.CounterServerImpl{})
	listener, err := net.Listen("tcp", ":8089")
	if err != nil {
		panic(err)
	}
	err = server.Serve(listener)
	if err != nil {
		panic(err)
	}
}
