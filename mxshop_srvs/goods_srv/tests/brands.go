package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	"mxshop_srvs/goods_srv/proto"
)

var brandClient proto.GoodsClient
var conn *grpc.ClientConn


func TestGetBrandList(){
	rsp, err := brandClient.BrandList(context.Background(), &proto.BrandFilterRequest{
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Total)
	for _, brand := range rsp.Data {
		fmt.Println(brand.Name)
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
	brandClient = proto.NewGoodsClient(conn)
}

func main() {
	Init()
	TestGetBrandList()

	conn.Close()
}