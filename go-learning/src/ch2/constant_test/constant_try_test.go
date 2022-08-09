package constant_test

import "testing"

const (
	Monday = iota + 1
	Tuesday
	Wednesday
)

// 用于对状态位的标识
const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestConstantTry(t *testing.T) {
	t.Log(Monday, Tuesday)
}
func TestConstantTry1(t *testing.T) {
	a := 1
	t.Log(a&Readable, a&Writable, a&Executable)
}
