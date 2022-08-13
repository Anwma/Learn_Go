package main

func main() {
	i := 0
Here: //这行的第一个词，以冒号结束作为标签
	println(i)
	i++
	if i == 10{
		return
	}
	goto Here //跳转到Here去
}
