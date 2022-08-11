// 7.13
package main

import (
	"unsafe"
)

/*
func main() {
	str := "hello,world"

	sh := *(*reflect.StringHeader)(unsafe.Pointer(&str))
	page := getPage(uintptr(unsafe.Pointer(sh.Data)))
	//Linux 中
	//syscall.Mprotect(page, syscall.PROT_READ|syscall.PROT_WRITE) // 改变内存页的只读权限

	fmt.Println("原始字符串：", str)
	bytes := (*(*[0xFF]byte)(unsafe.Pointer(sh.Data)))[:len(str)]
	fmt.Println("反转字符串：", bytes2string(reverse(bytes)))
}

func reverse(s []byte) []byte {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func getPage(p uintptr) []byte {
	return (*(*[0xFFFFFF]byte)(unsafe.Pointer(p & ^uintptr(syscall.Getpagesize()-1))))[:syscall.Getpagesize()]
}

func bytes2string(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

*/

func string2bytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}
