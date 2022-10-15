package main

import "fmt"

func main() {
	fmt.Println("Welcome to the class on poniters")

	// var ptr *int
	// fmt.Println("value of pointer is : ", ptr)
	myNumber := 55
	var ptr = &myNumber

	fmt.Println("Value of pointer is ", myNumber)
	fmt.Println("Value of pointer address is ", ptr)

	*ptr = *ptr * 2
	fmt.Println("New value: ", myNumber)

}
