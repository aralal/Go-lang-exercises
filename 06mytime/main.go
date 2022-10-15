package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createddate := time.Date(2020, time.May, 21, 23, 34, 0, 0, time.UTC)
	fmt.Println(createddate)
	fmt.Println(createddate.Format("01-02-2006 Monday"))

}
