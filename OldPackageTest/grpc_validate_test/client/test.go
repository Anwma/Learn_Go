package main

func main(){
	a := 1.01
	for i := 0; i < 1; i++ {
		a *= a
	}
	println(a)
}