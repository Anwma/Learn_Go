package mycalc

import . "fmt"
var name string

func init()  {
	Println("init了sub")

}
func Subtract(x, y int) int {
	return x - y
}
