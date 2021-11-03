package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"time"
)

func main() {
	c, _ := rocketmq.NewPushConsumer(
		consumer.WithNameServer([]string{"192.168.199.130:9876"}),
		consumer.WithGroupName("mxshop"),

	)
<<<<<<< HEAD
	if err := c.Subscribe("TransTopic", consumer.MessageSelector{}, func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
=======
	if err := c.Subscribe("TransTopic123", consumer.MessageSelector{}, func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
>>>>>>> 46eb3b74e18e70cbe7738bdbe69f4a5cf2a72cb6
		for i := range msgs {
			fmt.Printf("GGGGGGGGGGet  msg :%v\n", msgs[i])
		}
		return consumer.ConsumeSuccess, nil
	}); err != nil {
		fmt.Println("get msg fail!")
	}
	_ = c.Start()
	time.Sleep(time.Hour)
	_ = c.Shutdown()
}
