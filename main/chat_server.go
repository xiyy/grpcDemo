package main

import (
	"google.golang.org/grpc"
	"grpcDemo/chat"
	"log"
	"net"
)

func main() {
	log.Println("启动服务端...")
	server := grpc.NewServer()
	// 注册 ChatServer
	chat.RegisterChatServer(server, &chat.ChatServerImpl{})
	address, err := net.Listen("tcp", ":3000")
	if err != nil {
		panic(err)
	}
	if err := server.Serve(address); err != nil {
		panic(err)
	}
}
