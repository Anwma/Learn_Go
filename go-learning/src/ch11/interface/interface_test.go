package interface_test

import "testing"

type Programmer interface {
	WriteHelloWorld() string
}

type GoProgrammer struct {
}

func (g *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"Hello World\")"
}

func TestClient(t *testing.T) {
	//p是类型为Programmer的接口变量
	var p Programmer
	//p内部的数据为实现GoProgrammer的一个实例
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())
}
