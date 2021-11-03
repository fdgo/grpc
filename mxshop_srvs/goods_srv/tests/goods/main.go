package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"mxshop_srvs/goods_srv/proto"
)


var brandClient proto.GoodsClient
var conn *grpc.ClientConn

func TestGetGoodsList(){
	rsp, err := brandClient.GoodsList(context.Background(), &proto.GoodsFilterRequest{
		TopCategory: 130361,
		PriceMin: 90,
		//KeyWords: "深海速冻",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Total)
	for _, good := range rsp.Data {
		fmt.Println(good.Name, good.ShopPrice)
	}
}

func TestBatchGetGoods(){
	rsp, err := brandClient.BatchGetGoods(context.Background(), &proto.BatchGoodsIdInfo{
		Id: []int32{421, 422, 423},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Total)
	for _, good := range rsp.Data {
		fmt.Println(good.Name, good.ShopPrice)
	}
}

func TestGetGoodsDetail(){
	rsp, err := brandClient.GetGoodsDetail(context.Background(), &proto.GoodInfoRequest{
		Id: 421,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Name)
	fmt.Println(rsp.DescImages)
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
	//TestCreateUser()
	//TestGetGoodsList()
	//TestBatchGetGoods()
	TestGetGoodsDetail()

	conn.Close()
}
