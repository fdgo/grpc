package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"time"
)

type OrderListener struct {

}

func (o *OrderListener)ExecuteLocalTransaction(msg * primitive.Message)primitive.LocalTransactionState  {
	fmt.Println("开始执行本地逻辑")
	time.Sleep(time.Second*3)
	fmt.Println("执行本地逻辑fail")
	return primitive.UnknowState
}
func (o *OrderListener)CheckLocalTransaction(msg * primitive.MessageExt)primitive.LocalTransactionState  {
	fmt.Println("执行消息回查")
	return primitive.CommitMessageState
}
func main()  {
	p, err := rocketmq.NewTransactionProducer(
		&OrderListener{},
		producer.WithNameServer([]string{"192.168.164.128:9876"}),)
	if err != nil {
		panic("create producer fail!")
	}
	if err = p.Start(); err != nil {
		panic("start producer fail!")
	}
	res,err := p.SendMessageInTransaction(context.Background(),primitive.NewMessage("TransTopic",[]byte("this is transaction message 777777777777")))
	if err != nil {
		fmt.Println("producer fail:%s\n", err.Error())
	} else {
		fmt.Printf("producer ok:%s", res.String())
	}
	//res.State
	time.Sleep(time.Hour)
	_ = p.Shutdown()
}