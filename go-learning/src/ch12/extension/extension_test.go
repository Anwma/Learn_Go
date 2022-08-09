package extension

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Print("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

type Dog struct {
	//p *Pet
	Pet
}

//func (d *Dog) Speak() {
//	fmt.Print("Wang!")
//}
//func (d *Dog) SpeakTo(host string) {
//	d.Speak()
//	fmt.Println(" ", host)
//}

func (d *Dog) Speak() {
	fmt.Print("Wang!")
}

func TestDog(t *testing.T) {
	//cannot use new(Dog) (value of type *Dog) as type Pet in variable declaration
	//var dog Pet = new(Dog) //不支持LSP原则

	/*
		内嵌的结构类型无法当作继承来使用，不支持访问子类的方法、数据
		也就是不支持重载，不支持LSP。
	 */
	dog := new(Dog)
	//dog.Speak()
	dog.SpeakTo("Chao")
}
