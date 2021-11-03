package main

import (
	"context"
	"fmt"
	"time"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

type OrderListener struct{}

func (o *OrderListener) ExecuteLocalTransaction(msg *primitive.Message) primitive.LocalTransactionState {
	fmt.Println("开始执行本地逻辑")
<<<<<<< HEAD
	time.Sleep(time.Second * 3)
	fmt.Println("执行本地逻辑失败")
	//本地执行逻辑无缘无故失败 代码异常 宕机
	return primitive.UnknowState //fmt.Println("执行本地逻辑成功") primitive.CommitMessageState
=======
	time.Sleep(time.Second*3)
	fmt.Println("执行本地逻辑失败")
	//本地执行逻辑无缘无故失败 代码异常 宕机
	return primitive.UnknowState
>>>>>>> 46eb3b74e18e70cbe7738bdbe69f4a5cf2a72cb6
}

func (o *OrderListener) CheckLocalTransaction(msg *primitive.MessageExt) primitive.LocalTransactionState {
	fmt.Println("rocketmq的消息回查")
<<<<<<< HEAD
	time.Sleep(time.Second * 15)
=======
	time.Sleep(time.Second*15)
>>>>>>> 46eb3b74e18e70cbe7738bdbe69f4a5cf2a72cb6
	return primitive.CommitMessageState
}

func main() {
	p, err := rocketmq.NewTransactionProducer(
		&OrderListener{},
<<<<<<< HEAD
		producer.WithNameServer([]string{"192.168.199.131:9876"}),
	)
=======
		producer.WithNameServer([]string{"192.168.0.104:9876"}),
		)
>>>>>>> 46eb3b74e18e70cbe7738bdbe69f4a5cf2a72cb6
	if err != nil {
		panic("生成producer失败")
	}

<<<<<<< HEAD
	if err = p.Start(); err != nil {
		panic("启动producer失败")
	}
=======
	if err = p.Start(); err != nil {panic("启动producer失败")}
>>>>>>> 46eb3b74e18e70cbe7738bdbe69f4a5cf2a72cb6

	res, err := p.SendMessageInTransaction(context.Background(), primitive.NewMessage("TransTopic", []byte("this is transaction message2")))
	if err != nil {
		fmt.Printf("发送失败: %s\n", err)
<<<<<<< HEAD
	} else {
		fmt.Printf("发送成功: %s\n", res.String())
	}

	time.Sleep(time.Hour)
	if err = p.Shutdown(); err != nil {
		panic("关闭producer失败")
	}
=======
	}else{
		fmt.Printf("发送成功: %s\n", res.String())
	}
	
	time.Sleep(time.Hour)
	if err = p.Shutdown(); err != nil {panic("关闭producer失败")}
>>>>>>> 46eb3b74e18e70cbe7738bdbe69f4a5cf2a72cb6
}
