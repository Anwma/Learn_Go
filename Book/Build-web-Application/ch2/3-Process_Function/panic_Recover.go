package main

import (
	"fmt"
	"os"
)

/*

	import(
		. "fmt"
	)

	import(
		f "fmt"
	)

*/
/*

	import (
		"database/sql"
		_ "github.com/ziutek/mymysql/godrv"
	)

*/
func main() {
	fmt.Println("hello world")
}

func init() {
	var user = os.Getenv("USER")

	if user == "" {
		panic("no value for $USER")
	}
}

func throwsPanic(f func()) (b bool) {
	defer func() {
		if x := recover(); x != nil {
			b = true
		}
	}()
	f() //执行函数f，如果f中出现了panic，那么就可以恢复回来
	return
}
