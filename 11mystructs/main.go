package main

import "fmt"

func main() {
	fmt.Println("Welcome to structs in golang")
	//no inheritance in golang, No super or parent
	aravind := User{"Aravind", 20, "aravind@google", true}
	fmt.Println(aravind)
	fmt.Printf("aravind details are: %+v \n", aravind)
}

type User struct {
	Name   string
	Age    int
	Email  string
	Status bool
}
