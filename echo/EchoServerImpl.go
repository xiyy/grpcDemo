package echo

import (
	"context"
	"fmt"
)

type EchoServerImpl struct {
}

//实现EchoServer服务Echo(context.Context, *EchoReq) (*EchoRes, error),在该方法中处理业务逻辑
func (echoServer *EchoServerImpl) Echo(context context.Context, echoReq *EchoReq) (*EchoRes, error) {
	fmt.Printf("message from client: %v\n", echoReq.GetMsg())
	resp := &EchoRes{Msg: echoReq.GetMsg()}
	return resp, nil
}
