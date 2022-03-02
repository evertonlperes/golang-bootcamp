package main

import "fmt"

func main() {

	// approach #1
	// var colors map[string]string

	// approach #2
	colors := make(map[string]string)

	// approach #3
	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#00ff00",
	// 	"blue":  "#0000ff",
	// }

	colors["white"] = "#ffffff"

	// delete a map
	delete(colors, "white")
	fmt.Println(colors)
}
