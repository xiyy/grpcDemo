package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpcDemo/counter"
	"os"
	"strconv"
)

func main() {
	start, _ := strconv.ParseInt(os.Args[1], 10, 64)

	clientConn, err := grpc.Dial("127.0.0.1:8089", grpc.WithInsecure())
	if err != nil {
		fmt.Errorf("dial failed. err: [%v]\n", err)
		return
	}
	//客户端的 Count 接口返回的是一个 stream，不断地调用这个 stream 的 Recv 方法，可以不断地获取来自服务端的返回
	client := counter.NewCounterClient(clientConn)
	stream, err := client.Count(context.Background(), &counter.CountReq{
		Start: start,
	})
	if err != nil {
		fmt.Errorf("count failed. err: [%v]\n", err)
		return
	}

	for {
		res, err := stream.Recv()
		if err != nil {
			fmt.Errorf("client count failed. err: [%v]", err)
			return
		}

		fmt.Printf("server count: %v\n", res.GetNum())
	}
}
