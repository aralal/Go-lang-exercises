package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user input!"
	fmt.Println(welcome)

	// reading things
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for the site:")

	// coma ok || coma err   (coma error syntax)

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for the rating,", input)
	fmt.Printf("Type of this rating is %T \n", input)

	reader1 := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the age:")
	input1, _ := reader1.ReadString('\n')
	fmt.Printf("you are %v years old .", input1)
}
