package main

import "fmt"

func main() {
	fmt.Println("Welcome to structs in golang")
	//no inheritance in golang, No super or parent
	aravind := User{"Aravind", 20, "aravind@google", true}
	fmt.Println(aravind)
	fmt.Printf("aravind details are: %+v \n", aravind)
	aravind.GetStatus()
	aravind.Newmail()
	fmt.Println("Email is :", aravind.Email)
}

type User struct {
	Name   string
	Age    int
	Email  string
	Status bool
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}
func (u User) Newmail() {
	u.Email = "test@go.dev"
	fmt.Println("New mail :", u.Email)
}
