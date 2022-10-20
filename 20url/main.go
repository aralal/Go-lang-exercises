package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://loc.dev:3000/learn?coursename=reactjs&paymentid=hshdfkjsh"

func main() {
	fmt.Println("Welcome to handling urls in golang")
	fmt.Println(myurl)

	result, err := url.Parse(myurl)
	checkNilerr(err)

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Port())
	// fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	qparams := result.Query()
	fmt.Printf("The type of query params are :%T \n", qparams)

	fmt.Println(qparams["coursename"])
	fmt.Println(qparams["paymentid"])

	for _, val := range qparams {
		fmt.Println("Param is:", val)
	}

	partsofurl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}

	anotherUrl := partsofurl.String()
	fmt.Println(anotherUrl)

}

func checkNilerr(err error) {
	if err != nil {
		panic(err)
	}
}
