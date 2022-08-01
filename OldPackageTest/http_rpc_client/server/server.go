package main

import (
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
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
	_ = rpc.RegisterName("HelloService", &HelloService{})
	http.HandleFunc("/jsonrpc", func(writer http.ResponseWriter, request *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			ReadCloser: request.Body,
			Writer:     writer,
		}
		err := rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
		if err != nil {
			return 
		}
	})
	err := http.ListenAndServe(":1234",nil)
	if err != nil {
		return 
	}
}
