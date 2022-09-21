package main

import "fmt"

type alive interface {
	walk()
}
type People struct {
	name string
	tel  string
}

func (p People) walk()  {
	
}
func main() {
	fmt.Print()
}
