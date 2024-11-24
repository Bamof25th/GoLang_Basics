package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println("Number is :", i)
	}

	counter := 0
	for {
		fmt.Println("infinite loop")
		counter++

		if counter == 3 {
			break
		}
	}
	numbers := []int{1, 2, 3, 4, 5, 6}
	for index, value := range numbers {
		fmt.Printf("the index of number is %d, and the number is %d\n", index, value)
	}
}
