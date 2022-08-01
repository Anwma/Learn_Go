package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	//1.建立连接
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		panic("连接失败")
	}
	//var reply *string = new(string)
	var reply string
	err2 := client.Call("HelloService.Hello", "jeff", &reply)
	if err2 != nil {
		panic("调用失败")
	}
	fmt.Println(reply) //0xc00004a0f0

}
