package main

import "fmt"

func main() {
	defer fmt.Println("World")
	defer fmt.Println("one")
	defer fmt.Println("two")

	fmt.Println("Hello")
	myDefer()
}

// world,one,two
// 0 ,1 ,2, 3 ,4
// hello,43210two,one,world
func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
