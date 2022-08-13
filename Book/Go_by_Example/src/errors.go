package main

import (
	"errors"
	"fmt"
)

//参数arg 返回int error两个参数
func f1(arg int) (int, error) {
	//若arg==42
	if arg == 42 {
		//返回int=-1 error 错误信息 can't work with 42
		return -1, errors.New("can't work with 42")

	}
	//若不是则返回arg+3 , nil
	return arg + 3, nil
}

//定义argError结构体
type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {

		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {
	//遍历{7,42}这个切片,空白化(索引值index)(key)值为_,value值为i。
	for _, i := range []int{7, 42} {
		//第一遍
		//i=7, r,e传入f1中的int error两个参数，当e 不为nil时
		//转到f1判断,因为i->arg=7 != 42 ->return arg+3,nil ->r=10,e=nil
		//不满足if条件 走else分支
		//第二遍
		//i=42,其他同上,向f1方法传入i=42的值->return -1,errors.New(can't work with 42)
		//满足if条件。打印f1 failed:can't work with 42------------------②
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r) //打印 f1 worked: 10 ------①
		}
	}
	//第一遍同上
	//第二遍 return -1, &argError{arg, "can't work with it"}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e) //打印 f2 failed: 42 - can't work with it----④
		} else {
			fmt.Println("f2 worked:", r) //打印 f2 worked: 10 ------③
		}
	}
	//e接收f2返回的信息 &argError{arg, "can't work with it"}
	_, e := f2(42)
	//ar ok 分别接收e的arg 和 字符串;当ok也就是字符串不为空时
	if ae, ok := e.(*argError); ok {
		//打印两条语句
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
