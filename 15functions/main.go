package main

import "fmt"

func main() {
	fmt.Println("Functions in golang")
	result := adder(3, 5)
	fmt.Println("The result is :", result)

	proRes, mymessage := proAdder(3, 45, 6, 3, 5)
	fmt.Println("The proresult is:", proRes)
	fmt.Println("The message is:", mymessage)

}
func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}
func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "hi"
}
