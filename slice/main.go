package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println("Learning about slices in GoLang")

	// numbers := []int{1, 23, 4, 5, 5, 6, 7}
	// numbers = append(numbers, 3, 4, 5, 6, 8, 79, 45, 2, 12, 23)
	// fmt.Println("Here is your slice:", numbers)
	// fmt.Printf("Here is your slice with data type of : %T\n", numbers)
	// fmt.Println("Here is length of your slice:", len(numbers))

	numbers := make([]int, 5, 6)
	numbers = append(numbers, 2, 1, 3, 4, 5, 6, 78)
	arr := [5]int{1, 2, 3, 4, 5}
	s := arr[1:4] // Slice contains elements 2, 3, 4

	fmt.Println("here is your slice :", numbers)
	fmt.Println("here is your slice :", s)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)
	// The slices package contains a number of useful utility functions for slices.

	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	// 2D slice
	twoD := make([][]int, 3)

	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
