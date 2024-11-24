package main

import "fmt"

func main() {
	x := 5
	if x > 4 {
		fmt.Println("x is greater than four")
	} else {
		fmt.Println("x is smaller than four")
	}

	z := 9
	if z > 10 {
		fmt.Println("z is greater than 10")
	} else if z > 15 {
		fmt.Println("z is greater than 10 and smaller than 15")
	} else {
		fmt.Println("z is smaller than 10 and 15")
	}
}
