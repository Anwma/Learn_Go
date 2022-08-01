package main

import (
	proto_bak "OldPackageTest/grpc_proto_test/proto-bak"
	"context"
	"fmt"
	"google.golang.org/grpc"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"time"
)

func main() {
	//创建连接
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Printf("连接失败：[%v]\n", err)
		return
	}
	defer conn.Close()
	//声明客户端
	client := proto_bak.NewGreeterClient(conn)
	rsp, _ := client.SayHello(context.Background(), &proto_bak.HelloRequest{
		Name: "anwma",
		Url:  "https://yuque.com",
		G:    proto_bak.Gender_MALE,
		Mp: map[string]string{
			"name":    "Golang",
			"company": "Jetbrains",
		},
		AddTime: timestamppb.New(time.Now()),
	})
	fmt.Println(rsp.Message)
}
