package main

import (
	"OldPackageTest/stream_grpc_test/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"sync"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	//服务端流模式
	c := proto.NewGreeterClient(conn)
	res, _ := c.GetStream(context.Background(), &proto.StreamReqData{Data: "是"})
	for {
		a, err2 := res.Recv() // 如果大家懂socket编程的话就明白 send recv
		if err2 != nil {
			fmt.Println(err2)
			break
		}
		fmt.Println(a.Data)
	}
	//客户端流模式
	putStr, _ := c.PutStream(context.Background())
	i := 0
	for {
		i++
		_ = putStr.Send(&proto.StreamReqData{
			Data: fmt.Sprintf("Go语言%d", i),
		})
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}

	//双向流模式
	allStr, _ := c.AllStream(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			data, _ := allStr.Recv()
			fmt.Println("收到客户端消息：" + data.Data)
		}
	}()
	//1.集中学习protobuf,grpc
	go func() {
		defer wg.Done()
		for {
			_ = allStr.Send(&proto.StreamReqData{
				Data: "hello,I am client",
			})
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()
}
