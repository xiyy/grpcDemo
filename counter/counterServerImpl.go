package counter

import (
	"fmt"
	"time"
)

type CounterServerImpl struct {
}

//实现CounterServer接口的Count(*CountReq, Counter_CountServer) error
func (countServer *CounterServerImpl) Count(countReq *CountReq, stream Counter_CountServer) error {
	fmt.Printf("request from client,start:%v", countReq.GetStart())
	start := countReq.GetStart()
	for {
		start++
		if start > 100 {
			break
		}
		err := stream.Send(&CountRes{Num: start})
		if err != nil {
			return err
		}
		time.Sleep(time.Millisecond * 500)
	}
	return nil

}
