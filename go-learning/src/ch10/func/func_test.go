package fn_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

//输入参数和返回参数均为函数类型
func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestFn(t *testing.T) {
	//需要引入随机数种子...
	a, _ := returnMultiValues()
	t.Log(a)
	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))
}

func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func TestVarParam(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4))
	t.Log(Sum(1, 2, 3, 4, 5))
}

func Clear() {
	fmt.Println("Clear resources.")
}

func TestDefer(t *testing.T) {
	//可将Clear()转化为匿名函数嵌套到其中
	//defer func() {
	//	fmt.Println("Clear resources.")
	//}()
	defer Clear()
	fmt.Println("Start")
	panic("err")
	//不可达
	//fmt.Println("End")
}