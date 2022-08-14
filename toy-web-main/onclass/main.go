package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是主页")
}

func user(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是用户")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是创建用户")
}

func order(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是订单")
}
func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是测试")
}

func main() {
	//server := NewHttpServer("test-server")
	//server.Route("/user/signup", SignUp)
	//http.HandleFunc("/", home)
	//http.HandleFunc("/user", user)
	//http.HandleFunc("/user/create", createUser)
	//http.HandleFunc("/order", order)
	//http.HandleFunc("/test", test)
	//http.ListenAndServe(":8080", nil)
}

//type Server interface {
//	Route(pattern string, handlerFunc http.HandlerFunc)
//	Start(address string) error
//}
//
//type sdkHttpServer struct {
//	Name string
//}
