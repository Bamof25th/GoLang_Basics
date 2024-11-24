package main

import (
	"fmt"
)


// func divide(a, b float64) (float64, error) {
// 	if b == 0 {
// 		return 0, fmt.Errorf("denominator must not be zero")
// 	}
// 	return a / b, nil
// }

func divide(a, b float64) (float64, string) {
	if b == 0 {
		return 0, "error"
	}
	return a / b, "nil"
}

func main() {
	fmt.Println("started Error handling in the functions")
	// ans, err := divide(10, 0)
	// if err != nil {
	// 	fmt.Println("error handling")
	// }
	// fmt.Println("division of two numbers is:", ans)
	ans, _ := divide(10, 0)
	fmt.Println("division of two numbers is:", ans)
}
