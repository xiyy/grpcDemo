package main

import (
	"google.golang.org/grpc"
	"grpcDemo/echo"
	"net"
)

func main() {
	server := grpc.NewServer()
	echo.RegisterEchoServer(server, &echo.EchoServerImpl{})
	listener, err := net.Listen("tcp", ":8088")
	if err != nil {
		panic(err)
	}
	err = server.Serve(listener)
	if err != nil {
		panic(err)
	}
}
