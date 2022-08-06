package main

import (
	"fmt"
)

func main()  {
	str := "user-srv"

	s := fmt.Sprintf(`Service == "%s"`,str)

	//`Service == "user-srv"`
	fmt.Println(s)
}
