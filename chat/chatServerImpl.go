package chat

import (
	"io"
	"log"
	"strconv"
)

type ChatServerImpl struct {
}

//实现ChatServer接口的BidStream(Chat_BidStreamServer) error 方法
func (chatServerImpl *ChatServerImpl) BidStream(stream Chat_BidStreamServer) error {
	ctx := stream.Context()
	for {
		select {
		case <-ctx.Done():
			log.Println("收到客户端通过context发出的终止信号")
			return ctx.Err()
		default:
			// 接收从客户端发来的消息
			input, err := stream.Recv()
			if err == io.EOF {
				log.Println("客户端发送的数据流结束")
				return nil
			}
			if err != nil {
				log.Println("接收数据出错:", err)
				return err
			}
			switch input.Input {
			case "end":
				log.Println("收到'结束对话'指令")
				if err := stream.Send(&Response{Output: "收到结束指令"}); err != nil {
					return err
				}
				// 收到结束指令时，通过 return nil 终止双向数据流
				return nil
			case "return data":
				log.Println("收到'返回数据流'指令")
				// 收到 收到'返回数据流'指令， 连续返回 10 条数据
				for i := 0; i < 10; i++ {
					if err := stream.Send(&Response{Output: "数据流 #" + strconv.Itoa(i)}); err != nil {
						return err
					}
				}
			default:
				// 缺省情况下， 返回 '服务端返回: ' + 输入信息
				log.Printf("[收到消息]: %s", input.Input)
				if err := stream.Send(&Response{Output: "服务端返回: " + input.Input}); err != nil {
					return err
				}
			}
		}
	}
}
