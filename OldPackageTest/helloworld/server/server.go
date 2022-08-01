package main

import (
	"net"
	"net/rpc"
)

type HelloService struct {
}

func (s *HelloService) Hello(request string, reply *string) error {
	//返回值是通过修改reply的值
	*reply = "hello, " + request
	return nil
}
func main() {
	//1. 实例化一个server
	listener, err := net.Listen("tcp",":1234")
	if err != nil {
	    return
	}
	//2. 注册处理逻辑
	err2 := rpc.RegisterName("HelloService", &HelloService{})
	if err2 != nil {
		return
	}
	//3. 启动服务
	conn, err3 := listener.Accept() //当一个新的连接进来的时候,
	if err3 != nil{
		return
	}
	rpc.ServeConn(conn)

	//一连串的代码大部分都是net的包好像和rpc没关系
	//不想 rpc调用中有几个问题需要解决 1.call id 2.序列化和反序列化
	//python下的开发而言 这个就显得不好用
	//可以跨语言调用 1.go语言的rpc的序列化协议是什么(Gob编码 go特有) 2.能否替换成常见的序列化
}
