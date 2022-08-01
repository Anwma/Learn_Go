package main

import (
	"fmt"
	"net/http"
)

type MyMux struct {
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayhelloName(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello myroute!")
}

func main() {
	mux := &MyMux{}
	http.ListenAndServe(":9090", mux)
}

/*

Go代码的执行流程
通过对http包的分析之后，现在让我们来梳理一下整个的代码执行过程。

首先调用Http.HandleFunc

按顺序做了几件事：

1 调用了DefaultServeMux的HandleFunc

2 调用了DefaultServeMux的Handle

3 往DefaultServeMux的map[string]muxEntry中增加对应的handler和路由规则

其次调用http.ListenAndServe(":9090", nil)

按顺序做了几件事情：

1 实例化Server

2 调用Server的ListenAndServe()

3 调用net.Listen("tcp", addr)监听端口

4 启动一个for循环，在循环体中Accept请求

5 对每个请求实例化一个Conn，并且开启一个goroutine为这个请求进行服务go c.serve()

6 读取每个请求的内容w, err := c.readRequest()

7 判断handler是否为空，如果没有设置handler（这个例子就没有设置handler），handler就设置为DefaultServeMux

8 调用handler的ServeHttp

9 在这个例子中，下面就进入到DefaultServeMux.ServeHttp

10 根据request选择handler，并且进入到这个handler的ServeHTTP

  mux.handler(r).ServeHTTP(w, r)
11 选择handler：

A 判断是否有路由能满足这个request（循环遍历ServeMux的muxEntry）
B 如果有路由满足，调用这个路由handler的ServeHTTP
C 如果没有路由满足，调用NotFoundHandler的ServeHTTP

*/

//Conn 的 goroutine
/*
	Go在等待客户端请求里是这样写的

		c, err := srv.newConn(rw)
		if err != nil {
		    continue
		}
		go c.serve()

	客户端的每次请求都会创建一个Conn 这个Conn里面保存了该次请求的
	信息，然后再传递道对应的handler，该handler 中便可读取到相应的
	header信息，这样就保证了每个请求的独立性
*/
/*

//ServeMux 的自定义
type ServeMux struct {
	mu    sync.RWMutex        //锁，由于请求涉及到并发处理，因此这里需要一个锁机制
	m     map[string]muxEntry // 路由规则，一个string对应一个mux实体，这里的string就是注册的路由表达式
	hosts bool                // 是否在任意的规则中带有host信息
}
//muxEntry 注册的路由表达式
type muxEntry struct {
	explicit bool    // 是否精确匹配
	h        Handler // 这个路由表达式对应哪个handler
	pattern  string  //匹配字符串
}
//Handler 接口 =
type Handler interface {
	ServeHTTP(ResponseWriter, *Request) // 路由实现器
}

type HandlerFunc func(ResponseWriter, *Request)

// ServeHTTP calls f(w, r).
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}
//默认的路由器实现了 ServeHTTP
func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request) {
	//如果是*那么关闭链接，不然调用mux.Handler(r)返回对应设置路由
	//的处理Handler，然后执行h.ServeHTTP(w, r)
	if r.RequestURI == "*" {
		w.Header().Set("Connection", "close")
		w.WriteHeader(StatusBadRequest)
		return
	}
	h, _ := mux.Handler(r)
	h.ServeHTTP(w, r)
}

func (mux *ServeMux) Handler(r *Request) (h Handler, pattern string) {
	if r.Method != "CONNECT" {
		if p := cleanPath(r.URL.Path); p != r.URL.Path {
			_, pattern = mux.handler(r.Host, p)
			return RedirectHandler(p, StatusMovedPermanently), pattern
		}
	}
	return mux.handler(r.Host, r.URL.Path)
}

func (mux *ServeMux) handler(host, path string) (h Handler, pattern string) {
	mux.mu.RLock()
	defer mux.mu.RUnlock()

	// Host-specific pattern takes precedence over generic ones
	if mux.hosts {
		h, pattern = mux.match(host + path)
	}
	if h == nil {
		h, pattern = mux.match(path)
	}
	if h == nil {
		h, pattern = NotFoundHandler(), ""
	}
	return
}

*/
