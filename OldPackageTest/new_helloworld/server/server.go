package main

import (
	"OldPackageTest/new_helloworld/handler"
	"OldPackageTest/new_helloworld/server_proxy"
	"net"
	"net/rpc"
)

func main() {

	//1. 实例化一个server
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		return
	}
	//2. 注册处理逻辑 handler
	err2 := server_proxy.RegisterHelloService(&handler.NewHelloService{})
	if err2 != nil {
		return
	}
	//3. 启动服务
	for {
		conn, err3 := listener.Accept() //当一个新的连接进来的时候,
		if err3 != nil {
			return
		}
		go rpc.ServeConn(conn)
	}
}
