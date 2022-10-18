package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")
	content := "This need to go in the file - Hello there guys!"

	file, err := os.Create("./myfile.txt")
	// if err != nil {
	// 	panic(err)
	// }
	checkNilerr(err)

	length, err := io.WriteString(file, content)
	// if err != nil {
	// 	panic(err)
	// }
	checkNilerr(err)

	fmt.Println("length is:", length)
	defer file.Close()
	readFile("./myfile.txt")
}

func readFile(filename string) {
	//databyte, err := ioutil.ReadFile(filename)
	databyte, err := os.ReadFile(filename)
	checkNilerr(err)
	// if err != nil {
	// 	panic(err)
	// }

	fmt.Println("Text data inside the file is \n", string(databyte))
}

func checkNilerr(err error) {
	if err != nil {
		panic(err)
	}
}
