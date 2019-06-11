package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpcDemo/echo"
	"os"
	"strings"
)

func main() {
	clientConn, err := grpc.Dial("127.0.0.1:8088", grpc.WithInsecure())
	if err != nil {
		fmt.Errorf("dial failed.err:[%v]\n", err)
	}
	echoClient := echo.NewEchoClient(clientConn)
	response, err := echoClient.Echo(context.Background(), &echo.EchoReq{Msg: strings.Join(os.Args[1:], "")})
	if err != nil {
		fmt.Errorf("client echo failed,err:%v", err)
	}
	fmt.Printf("message from server:%v", response.GetMsg())

}
