package main

import "fmt"

func main() {
	fmt.Println("Basic usage:")
	basicMaps()
	fmt.Println("Different key value:")
	differentMap()
}

func basicMaps() {
	var colors map[string]string
	colors = make(map[string]string)
	colors["red"] = "#FF0000"
	colors["green"] = "#00FF00"
	colors["blue"] = "#0000FF"
	fmt.Println(colors)
}

func differentMap() {
	ages := map[string]int{
		"Numan": 23,
		"Ali":   24,
		"Veli":  19,
	}
	fmt.Println(ages)
}
