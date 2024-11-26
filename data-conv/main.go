package main

import (
	"fmt"
	"strconv"
)

func main() {

	var num int = 4

	fmt.Println("The number is: ", num)
	fmt.Printf("The type of the number is: %T\n", num)

	// num = num + 1.4

	var data float64 = float64(num)
	// data = data + 1.23
	fmt.Println("The number in float64 is: ", data)
	fmt.Printf("The type of the number in float64 is: %T\n", data)

	num = 12344

	str := strconv.Itoa(num)

	fmt.Println("The number in string is: ", str)
	fmt.Printf("The type of the number in string is: %T\n", str)

	number_string := "21.233"

	number_float, _ := strconv.ParseFloat(number_string, 64)

	fmt.Println("The number in float64 is: ", number_float)
	fmt.Printf("The type of the number in float64 is: %T\n", number_float)

}
