package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"time"
)

const (
	timestampFormat = time.StampNano //"Jan_2 15:04:05.000"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:50001", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := pb.NewGreeterClient(conn)
	md := metadata.Pairs("timestamp",time.Now().Format(timestampFormat))
	ctx := metadata.NewOutgoingContext(context.Background(),md)
	resp,err := client.SayHello(ctx,&pb.HelloRequest{Name:"hello,world"})
	if err == nil {
		fmt.Printf("Reply is %s\n",resp.Message)
	}else{
		fmt.Printf("call server error:%s\n",err)
	}
}
