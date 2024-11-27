package main

import "fmt"

func add(n int) int {
	return n * 20
}

func main() {
	// defer  its a key word similar to finally in javascript it make the line of code run at the very last

	fmt.Println("First")
	a := add(23)
	defer fmt.Println("After adding", a)
	defer fmt.Println("Middle")
	fmt.Println("Last")

    // LIFO  upar wala baadme chalega niche wala phele caue it gets initiated in stack way
	// defer fmt.Println("After adding", a)
	// defer fmt.Println("Middle")
}
