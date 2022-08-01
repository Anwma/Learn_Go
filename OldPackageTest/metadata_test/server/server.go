package main

import (
	"OldPackageTest/grpc_test/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"net"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {
	md,ok := metadata.FromIncomingContext(ctx)
	if ok{
		fmt.Println("get metadata error")
	}
	if nameSlice,ok:=md["name"];ok{
		fmt.Println(nameSlice)
		for i, e := range nameSlice {
			fmt.Println(i,e)
		}
	}
	//for key, val := range md{
	//	fmt.Println(key,val)
	//}
	return &proto.HelloReply{
		Message: "hello " + request.Name,
	}, nil
}

func main() {
	g := grpc.NewServer()
	proto.RegisterGreeterServer(g,&Server{})
	lis,err := net.Listen("tcp","0.0.0.0:8080")
	if err != nil {
	    panic("failed to listen" + err.Error())
	}
	err = g.Serve(lis)
	if err != nil {
		panic("failed to start" + err.Error())
	}
}
