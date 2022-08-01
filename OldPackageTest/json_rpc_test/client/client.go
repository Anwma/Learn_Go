package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	//1.建立连接
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		panic("连接失败")
	}
	//var reply *string = new(string)
	var reply string
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	err2 := client.Call("HelloService.Hello", "jeff", &reply) //底层是json字符串
	if err2 != nil {
		panic("调用失败")
	}
	fmt.Println(reply) //0xc00004a0f0

}
//{"method":"HelloService.Hello", "params":["hello"],"id":0}
