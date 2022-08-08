package main

import "fmt"

func main() {
	a := []int{7, 8, 9}
	fmt.Printf("%v", a[real(2)])
}

//func main()  {
//	ctx,_ := context.WithTimeout(context.Background(),0)
//	<-ctx.Done()
//	fmt.Println("timed out")
//}

func foo() *int {
	var i int
	i = 10
	return &i
}
