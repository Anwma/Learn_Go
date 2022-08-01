package main

import "fmt"

func main() {
	//seasons := []string{"Spring", "Summer", "Autumn", "Winter"}
	//for ix, season := range seasons {
	//	fmt.Printf("Season %d is: %s\n", ix, season)
	//}
	//
	//var season string
	//for _, season = range seasons {
	//	fmt.Printf("%s\n", season)
	//}
	items := [...]int{10, 20, 30, 40, 50}
	for i := 0; i < len(items); i++ {
		items[i] *= 2
	}
	//for _, item := range items {
	//	item *= 2
	//}
	fmt.Println(items)
}
