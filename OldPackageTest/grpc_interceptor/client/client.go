package main

import (
	"OldPackageTest/grpc_test/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"time"
)

func main() {
	//stream
	interceptor := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error{
		start := time.Now()
		err := invoker(ctx, method, req, reply, cc, opts...)
		if err != nil {
			return err
		}
		fmt.Printf("耗时: %s\n",time.Since(start))
		return err
	}
	var opts []grpc.DialOption
	opts = append(opts,grpc.WithInsecure())
	opts = append(opts,grpc.WithUnaryInterceptor(interceptor))
	//opt := grpc.WithUnaryInterceptor(interceptor)
	//conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure(),opt)
	conn, err := grpc.Dial("127.0.0.1:8080", opts...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := proto.NewGreeterClient(conn)
	r, err := c.SayHello(context.Background(), &proto.HelloRequest{Name: "Anwma"})
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Message)
}
