package util

import (
	"math"
	"strings"
)

// BinaryFormat 输出一个int2对应的二进制表示
func BinaryFormat(n int) string {
	sb := strings.Builder{}
	c := int(math.Pow(2, 31))
	//将n的二进制表示出来，写入到sb字符串当中并返回
	for i := 0; i < 32; i++ {
		if n&c != 0 {
			sb.WriteString("1")
		} else {
			sb.WriteString("0")
		}
		c >>= 1
	}
	return sb.String()
}
