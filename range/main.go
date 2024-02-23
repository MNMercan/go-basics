package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	for i, num := range numbers {
		fmt.Println("index:", i, "number:", num)
	}
	fmt.Println("-----------")
	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
		"blue":  "#0000FF",
	}
	for color, code := range colors {
		fmt.Println("Color:", color, " | Code: ", code)
	}
}
