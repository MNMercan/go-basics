package main

import "fmt"

func main() {
	var num, funcSelect int
	fmt.Println("Please set a number")
	fmt.Scan(&num)
	fmt.Println("Which func do you prefer?")
	fmt.Println("1 - Simple function")
	fmt.Println("2 - Comples function")
	fmt.Scan(&funcSelect)

	if funcSelect == 1 {
		simpleCon(num)
	} else {
		complesCon(num)
	}

}

func complesCon(num int) {
	if num > 10 && num < 20 {
		fmt.Println("Your number is between 10 and 20")
	} else if num > 20 && num < 30 {
		fmt.Println("Your number is between 20 and 30")
	} else {
		fmt.Println("x is not between 10 and 30")
	}
}

func simpleCon(x int) {
	if x > 0 {
		fmt.Println("Positive number")
	} else if x < 0 {
		fmt.Println("Negative number")
	} else {
		fmt.Println("Zero")
	}
}
