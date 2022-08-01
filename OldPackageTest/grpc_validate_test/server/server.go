package main

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net"

	"google.golang.org/grpc"

	"OldPackageTest/grpc_validate_test/proto"
)


type Server struct{}

func (s *Server) SayHello(ctx context.Context, request *proto.Person) (*proto.Person,
	error){
	return &proto.Person{
		Id: 32,
	}, nil
}

type Validator interface {
	Validate() error
}

func main(){
	//使用示例：
	//p := new(proto.Person)
	//p.Id = 1000
	//err := p.Validate()
	//if err != nil {
	//    panic(err)
	//}
	var interceptor grpc.UnaryServerInterceptor
	interceptor = func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		// 继续处理请求
		//(proto.Person) 可以满足person的验证 所有的接口 还有其他的接口 那怎么办呢?
		if r, ok := req.(Validator); ok {
		//if r, ok := req.(proto.Person); ok {
			if err := r.Validate(); err != nil {
				return nil, status.Error(codes.InvalidArgument, err.Error())
			}
		}

		return handler(ctx, req)
	}
	var opts []grpc.ServerOption
	opts = append(opts, grpc.UnaryInterceptor(interceptor))

	g := grpc.NewServer(opts...)
	proto.RegisterGreeterServer(g, &Server{})
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil{
		panic("failed to listen:"+err.Error())
	}
	err = g.Serve(lis)
	if err != nil{
		panic("failed to start grpc:"+err.Error())
	}



}