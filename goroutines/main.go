package main

import (
	"fmt"
	"time"
)

func flash(str string) {
	for i := 0; i < 3; i++ {
		fmt.Println(str, ":", i)
	}
}

func main() {

	// A goroutine is a lightweight thread of execution.
	fmt.Println("Starting goRoutines")

	flash("hello")

	// Create a new goroutine.
	go flash("world")

	//  unnamed function
	go func(i string) {
		fmt.Println(i)
	}("going")

	// Wait for 1 second to allow goroutines to finish.
	time.Sleep(time.Second)
	// fmt.Println("Main function finished")

	// Wait for  goroutine to finish.
	fmt.Println("All goroutines finished")


	

}


