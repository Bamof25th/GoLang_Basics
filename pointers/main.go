package main

import "fmt"

func modifyValue(a *int) {
	*a = *a * 20
}
func main() {
	fmt.Println("Hello, World!")

	var num int = 9

	ptr := &num

	// fmt.Printf("the number is : %d, and its pointer is %d\n", num, ptr)
	fmt.Printf("the number is : %d, and its pointer is %d\n", *ptr, ptr)

	var pointer *int

	if pointer == nil { // by default pointer has "nil" stored util its assigned with any variable
		fmt.Println("the pointer is nil")
	}
	value := 23
	modifyValue(&value)

	fmt.Println("the number is : ", value)

}
