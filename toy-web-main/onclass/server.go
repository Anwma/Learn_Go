package main

import (
	"fmt"
	"net/http"
)

type Server interface {
	// Route 设定一个路由 命中该路由会执行handlerFunc 的代码
	//method 可能是POST PUT GET
	Route(method string, pattern string, handlerFunc func(ctx *Context))
	// Start 启动我们的服务器
	Start(address string) error
}

// sdkHttpServer 这个是基于net/http 这个包实现的 http server
type sdkHttpServer struct {
	//Name server 的名字 给个标记 日志输出的时候用得上
	Name    string
	handler *HandlerBasedOnMap
}

// 我只要具备他的方法 那么我就实现了这个接口

// Route 注册路由
func (s *sdkHttpServer) Route(method string, pattern string, handlerFunc func(ctx *Context)) {
	//TODO implement me
	//http.HandleFunc(pattern, func(writer http.ResponseWriter,
	//	request *http.Request) {
	//	//if request.Method != method{
	//	//	writer.Write([]byte("error"))
	//	//}
	//	ctx := NewContext(writer,request)
	//	handlerFunc(ctx)
	//})
	key := s.handler.key(method,pattern)
	s.handler.handlers[key] = handlerFunc
}

func (s *sdkHttpServer) Start(address string) error {
	//TODO implement me
	http.Handle("/", s.handler)
	return http.ListenAndServe(address, nil)
}
func NewHttpServer(name string) Server {
	return &sdkHttpServer{
		Name: name,
	}
}

//上下文是依赖于框架自身来创建

func SignUp(ctx *Context) {
	req := &signUpReq{}

	err := ctx.ReadJson(req)
	if err != nil {
		ctx.BadRequestJson(err)
		return
	}

	resp := &commonResponse{
		Data: 123,
	}
	err = ctx.WriteJson(http.StatusOK, resp)

	//json.Marshal("confirmed_password") //过程式 命令式的写法

	//ctx.W.WriteHeader(err.Error())
	if err != nil {
		fmt.Printf("写入响应失败：%v", err)
	}
}

type signUpReq struct {
	Email             string `json:"email"`
	Password          string `json:"password"`
	ConfirmedPassword string `json:"confirmed_password"`
}

type commonResponse struct {
	//Tag 运行期间通过反射拿到 --声明式的API
	BizCode int         `json:"biz_code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

//type Header map[string][]string

//func NewServer() Server {
//	return &sdkHttpServer{}
//}

//func NewServerBaseOnXXX() Server {
//
//}

//type Factory func() Server
//
//var factory Factory
//
//func RegisterFactory(f Factory) {
//	factory = f
//
//}
//func NewServer() Server {
//	return factory()
//}
