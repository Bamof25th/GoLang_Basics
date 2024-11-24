package main

import "fmt"

func main() {
	age := 25
	name := "Alice"
	height := 5.4972316129

	// fmt.Println("Age :", age, "Height :", height, "Name :", name)
	// fmt.Printf("Age is %d\n ", age)               // format specifier default \n is for new line
	// fmt.Printf("Height is %.4f\n ", height)       // format specifier toFixed
	// fmt.Printf("Type of age is %T\n ", age)       // format specifier  typeOf
	// fmt.Printf("Type of height is %T\n ", height) // format specifier typeOf
	// fmt.Printf("Name is %s\n ", name)             // format specifier %s == string
	fmt.Printf("Name : %s Age : %d Height : %.4f ", name, age, height)
}
