package main

import (
	"OldPackageTest/stream_grpc_test/proto"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"sync"
	"time"
)

const PORT = "localhost:50052"

type server struct {
}

func (s *server) GetStream(req *proto.StreamReqData, res proto.Greeter_GetStreamServer) error {
	i := 0
	for {
		i++
		_ = res.Send(&proto.StreamResData{
			Data: fmt.Sprintf("%v", time.Now().Unix()),
		})
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}
	return nil
}
func (s *server) PutStream(cliStr proto.Greeter_PutStreamServer) error {
	for {
		if a, err := cliStr.Recv(); err != nil {
			//有错误的
			fmt.Println(err)
			break
		} else {
			fmt.Println(a.Data)
		}
	}
	return nil
}
func (s *server) AllStream(allStr proto.Greeter_AllStreamServer) error {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			data, _ := allStr.Recv()
			fmt.Println("收到客户端消息：" + data.Data)
		}
	}()
	go func() {
		defer wg.Done()
		for  {
			_ = allStr.Send(&proto.StreamResData{
				Data: "hello,I am server",
			})
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()
	return nil
}

func main() {
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &server{})
	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}
}
