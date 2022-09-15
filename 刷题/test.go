package main

func main() {
	m := make(map[int]int, 3)
	//初始化map为空
	x := len(m)
	m[1] = m[1]
	y := len(m)
	println(x, y)
}
