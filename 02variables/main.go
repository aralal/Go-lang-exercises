package main

import "fmt"

const LoggedinToken string = "bsdjsbkjdbk" //Public

func main() {
	var username string = "aravind"
	fmt.Println(username)
	fmt.Printf("This variable is of type: %T \n", username)

	var isloggedin bool = true
	fmt.Println(isloggedin)
	fmt.Printf("This variable is of type: %T \n", isloggedin)

	var smallvar uint8 = 255
	fmt.Println(smallvar)
	fmt.Printf("This variable is of type: %T \n", smallvar)

	var smallflot float64 = 255.364772883434
	fmt.Println(smallflot)
	fmt.Printf("This variable is of type: %T \n", smallflot)

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("this variable is of type: %T \n", anotherVariable)
	fmt.Println("Helllo")

	var website = "Learncodeonline.com"
	fmt.Println(website)

	noofusers := 40000.242
	fmt.Println(noofusers)

	fmt.Println(LoggedinToken)
	fmt.Printf("this variable is of type: %T \n", LoggedinToken)
}
