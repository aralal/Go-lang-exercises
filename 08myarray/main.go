package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")

	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "mango"
	fruitList[3] = "banana"

	fmt.Println("Frutis are :", fruitList)

	var vegList = [3]string{"potato", "mushroom", "beans"}
	fmt.Println("Veglist :", vegList)

}
