package main

import "fmt"

func main() {
	var arr [5]string

	arr[0] = "Aniket"
	arr[1] = "baghel"
	arr[2] = "Doe"
	arr[3] = "John"

	fmt.Println("numbers array :", arr)

	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers)
	fmt.Println(len(numbers))

	fmt.Println("2nd index contains", numbers[2])

}
