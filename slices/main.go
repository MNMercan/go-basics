package main

import "fmt"

// Dilin temel veri yapılarından biri olan dilimler (slices), Go programlama dilinde, boyutu değiştirilebilen,
//  esnek bir dizi veri yapısıdır. Dilimler, dizilerin bir alt kümesi olarak düşünülebilir ve Go dilinde sıkça kullanılır.

func main() {
	fmt.Println("Simple slice usage:")
	simpleSlice()
	fmt.Println("--------")
	fmt.Println("Slice usage with append:")
	appendSlice()
	fmt.Println("--------")
	fmt.Println("Subset Slice usage:")
	subsetSlice()
}

func simpleSlice() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(a)
}

func appendSlice() {
	var appendSlice []int
	appendSlice = append(appendSlice, 1, 2)
	fmt.Println(appendSlice)
	appendSlice = append(appendSlice, 3, 4, 5, 6)
	fmt.Println(appendSlice)
}

func subsetSlice() {
	mainSlice := []int{1, 2, 3, 4, 5}
	subsetSlice := mainSlice[1:4]
	fmt.Println("Main slice:", mainSlice)
	fmt.Println("Subset slice for mainSlice[1:4] :", subsetSlice)
}
