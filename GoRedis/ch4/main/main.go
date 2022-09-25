package main

import (
	"fmt"
	"unsafe"
)

type K struct {
}
type F struct {
	num1 K
	num2 int
}
type C struct {
	num2 F
}

func main() {
	s := "世界" //utf-8 编码
	for _, v := range s {
		fmt.Printf("%c\n", v)
	}
	//sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	//fmt.Println(sh.Len)
	//fmt.Println(unsafe.Sizeof("世界")) //string 底层 包含两个指针
	//fmt.Println(len("世界"))           //string 底层 包含两个指针
	//m := map[string]int{}       //map
	//m := map[string]struct{}{} //hashset
	//m["a"] = struct{}{}
	//a := make(chan struct{}) //只想发信号 不接收任何信息
	//如果只要键 不要值
	//a := C{}
	//b := K{}
	//b := int(0) //8
	//c := K{}
	//fmt.Println(unsafe.Sizeof(a))
	//a和c的地址一模一样
	//fmt.Printf("%p\n", &a) //0x110a5c0
	//fmt.Printf("%p\n", &b) //0x110a5c0
	//fmt.Printf("%p\n", &b)
	//fmt.Printf("%p\n", &c) //0x110a5c0
	//i := 1234
	//j := int32(1)
	//f := float32(3.141)
	//bytes := [5]byte{'h', 'e', 'l', 'l', 'o'}
	primes := [4]int{2, 3, 5}
	//p := &primes
	//
	//r := rune(666)
	//
	//fmt.Println(unsafe.Sizeof(i))      //int
	//fmt.Println(unsafe.Sizeof(j))      //int32
	//fmt.Println(unsafe.Sizeof(f))      //float32
	//fmt.Println(unsafe.Sizeof(bytes))  //byte
	fmt.Println(unsafe.Sizeof(primes)) //[]int len(arr) * type
	//fmt.Println(unsafe.Sizeof(p))      //ptr
	//fmt.Println(unsafe.Sizeof(r))      //rune
	//fmt.Println(unsafe.Sizeof("你好"))   //16??/

}
