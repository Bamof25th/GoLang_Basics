package main

import "fmt"

func main() {
	fmt.Println("Learning about slices in GoLang")

	numbers := []int{1, 23, 4, 5, 5, 6, 7}
	numbers = append(numbers, 3, 4, 5, 6, 8, 79, 45, 2, 12, 23)
	fmt.Println("Here is your slice:", numbers)
	fmt.Printf("Here is your slice with data type of : %T\n", numbers)
	fmt.Println("Here is length of your slice:", len(numbers))

}
