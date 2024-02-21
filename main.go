package main

import (
	"fmt"
	"strings"
)

func main() {
	var firstname, lastname string

	fmt.Println("Enter your firstname")
	fmt.Scanln(&firstname)
	fmt.Println("Enter your lastname")
	fmt.Scanln(&lastname)

	fullName := strings.TrimSpace(firstname + " " + lastname)

	if firstname == "" || lastname == "" {
		fmt.Println("Error: You should fill the firstname and lastname")
	} else {
		fmt.Println("Hello,", fullName)
	}
}
