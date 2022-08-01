package main

import "fmt"

func div(a, b int) (int, error) {
	if b == 0 {
		panic("被除数不能为0")
	}
	return a / b, nil
}
func main() {
	//错误就是能预知到可能出现的情况,这些情况会导致你的代码出问题 参数检查 数据库访问不了
	//data := 12
	//整数转字符串 字符串的内容形式不受限制
	//strconv.Itoa(data) //这个Itoa的函数不可能出错 所以没有必要返回error.
	//内部代码出错的时候, 应该抛出异常 panic
	//字符串转整数 因为字符串所包含的字符集是纯数字的超集 可能会出现中文字符 标点符号等情况 故会返回error
	//_, err := strconv.Atoi("12") //Atoi这个函数认为我的函数内部会出现一些预知错误情况
	//if err != nil {
	//	//错误
	//}
	//异常 go语言如何抛出异常和如何接收/接收异常
	//go语言认为错误就要自己处理,就个人而言,go语言的这种想法是正确的。但是实际使用中确实有点烦人
	a := 12
	b := 0
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("异常被捕获到")
		}
		fmt.Println("wozen")
	}()
	fmt.Println(div(a, b))

	//panic的坑

}
