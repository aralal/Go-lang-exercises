package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to json vedio")
	//EncodeJson()
	DecodeJson()

}

func EncodeJson() {
	lcocourses := []course{
		{"ReactJs Bootcamp", 299, "Learncodeonline.in", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 399, "Learncodeonline.in", "bcd123", []string{"fullstack", "js"}},
		{"Python Bootcamp", 499, "Learncodeonline.in", "efg123", nil},
	}

	//package this data as json data
	finalJson, err := json.MarshalIndent(lcocourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "MERN Bootcamp",
		"Price": 399,
		"website": "Learncodeonline.in",
		"tags": ["fullstack","js"]
	}
	`)
	var lcoCourse course
	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("Json was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v \n", lcoCourse)
	} else {
		fmt.Println("Json was not VAILD")
	}
	//for some cases you need to add key value pairs
	var myOnlinedata map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlinedata)
	fmt.Printf("%#v \n", myOnlinedata)

	for k, v := range myOnlinedata {
		fmt.Printf("The key %v and value is %v and type is %T \n", k, v, v)
	}

}
