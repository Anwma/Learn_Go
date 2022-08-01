package main

func main() {
	//var a byte
	//a = 'a' //输出ascii对应码值 。。 这里说明一下什么是ascii码
	//fmt.Println(a)
	//fmt.Printf("a=%c", a)
	a := 1.00
	for i := 0; i < 1; i++ {
		a *= 1.01
	}
	println(a)
}
