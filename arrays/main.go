package main

import "fmt"

func main() {
	var array1 [5]int
	
	array1[0] = 1
	array1[1] = 2
	array1[2] = 3
	array1[3] = 4
	array1[4] = 5
	fmt.Println(array1)
	//-----------
	array2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array2)

}
