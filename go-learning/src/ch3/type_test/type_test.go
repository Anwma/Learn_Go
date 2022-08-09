package type_test

import "testing"

type MyInt int64

func TestImplicit(t *testing.T) {
	var a int = 1
	var b int64
	b = int64(a)
	var c MyInt
	//各种隐式类型转换均不支持，即便是起一个别名
	c = MyInt(b)
	t.Log(a, b, c)
}
func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	//aPtr = aPtr + 1 //cannot convert 1 (untyped int constant) to *int
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}
func TestString(t *testing.T) {
	var s string
	t.Log("*" + s + "*")
	t.Log(len(s))
	//if s == "" {
	//
	//}
	
}
