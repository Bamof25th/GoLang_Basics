package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"isAdult"`
}

func main() {
	person := Person{Name: "John", Age: 23, IsAdult: true}

	fmt.Println(person)

	// convert to json encoding :: (Marshal JSON)

	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(jsonData))

	//  Decoding JSON data

	var newPerson Person
	err = json.Unmarshal(jsonData, &newPerson)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("person", newPerson)

}
