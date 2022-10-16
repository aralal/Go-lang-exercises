package main

import "fmt"

func main() {
	fmt.Println("If else in golang")

	logincount := 23
	var result string

	if logincount < 10 {
		result = "Regular user"
	} else {
		result = "something else"
	}
	fmt.Println(result)
}
