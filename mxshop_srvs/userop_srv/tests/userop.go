package main

import (
	"context"
	"google.golang.org/grpc"
	"mxshop_srvs/userop_srv/proto"
)


var userFavClient proto.UserFavClient
var messageClient proto.MessageClient
var addressClient proto.AddressClient
var conn *grpc.ClientConn

func TestAddressList(){
	_, err := addressClient.GetAddressList(context.Background(), &proto.AddressRequest{
		UserId: 1,
	})
	if err != nil {
		panic(err)
	}
}

func TestMessageList() {
	_, err := messageClient.MessageList(context.Background(), &proto.MessageRequest{
		UserId: 1,
	})
	if err != nil {
		panic(err)
	}
}

func TestUserFav() {
	_, err := userFavClient.GetFavList(context.Background(), &proto.UserFavRequest{
		UserId: 1,
	})
	if err != nil {
		panic(err)
	}
}


func Init(){
	var err error
<<<<<<< HEAD
	conn, err = grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
=======
	conn, err = grpc.Dial("192.168.164.128:50051", grpc.WithInsecure())
>>>>>>> 46eb3b74e18e70cbe7738bdbe69f4a5cf2a72cb6
	if err != nil {
		panic(err)
	}
	userFavClient = proto.NewUserFavClient(conn)
	messageClient = proto.NewMessageClient(conn)
	addressClient = proto.NewAddressClient(conn)
}



func main() {
	Init()

	TestAddressList()
	TestMessageList()
	TestUserFav()
	conn.Close()
}
