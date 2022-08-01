package handler

//为了解决命名冲突的问题

const HelloServiceName = "handler/HelloService"

//我们关心的是 NewHelloService 这个名字呢 还是这个结构体中的方法

type NewHelloService struct {
}

func (s *NewHelloService) Hello(request string, reply *string) error {
	//返回值是通过修改reply的值
	*reply = "hello, " + request
	return nil
}
