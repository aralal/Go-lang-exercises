package main

import "fmt"

func main() {

	fmt.Println("Maps in golang")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of languages:", languages)
	fmt.Println("JS short for:", languages["JS"])
	delete(languages, "RB")
	fmt.Println("List of languages:", languages)
	languages["RB"] = "ruby"
	languages["cpp"] = "C++"

	//loops in maps

	for key, value := range languages {
		fmt.Printf("For key %v , value is %v \n", key, value)
	}

}
