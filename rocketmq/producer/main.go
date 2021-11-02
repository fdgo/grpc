package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

func main() {
	p, err := rocketmq.NewProducer(producer.WithNameServer([]string{"192.168.199.130:9876"}))
	if err != nil {
		panic("create producer fail!")
	}
	if err = p.Start(); err != nil {
		panic("start producer fail!")
	}
	res, err := p.SendSync(context.Background(), primitive.NewMessage("imooc1", []byte("this is imooc1")))
	if err != nil {
		fmt.Println("producer fail:%s\n", err.Error())
	} else {
		fmt.Printf("producer ok:%s", res.String())
	}
	if err = p.Shutdown(); err != nil {
		panic("Shutdown fail!")
	}
}
