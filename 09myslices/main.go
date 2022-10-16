package main

import (
	"fmt"
)

func main() {
	// fmt.Println("Welcome to slices in golang")
	// var fruitList = []string{"apple", "orange", "peaches"}

	// fmt.Printf("type of fruit list is %T:", fruitList)

	// fruitList = append(fruitList, "mango", "banana")
	// fmt.Println(fruitList)

	// fruitList = append(fruitList[1:])
	// fmt.Println(fruitList)

	// highscores := make([]int, 4)

	// highscores[0] = 343
	// highscores[1] = 435
	// highscores[2] = 232
	// highscores[3] = 839

	// fmt.Println(highscores)
	// highscores = append(highscores, 564, 738, 666)
	// fmt.Println(highscores)
	// sort.Ints(highscores)
	// fmt.Println(highscores)
	// fmt.Println(sort.IntsAreSorted(highscores))

	var courses = []string{"java", "reactjs", "swift", "ruby", "python"}
	fmt.Println(courses)
	var index = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
