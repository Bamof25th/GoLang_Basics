package main

import (
	"fmt"
	"sync"
)

func worker(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("worker %d started\n", i)

	// some tasks are running
	//...
	fmt.Printf("worker %d finished\n", i)

}

func main() {
	fmt.Println("Exploring go routines")

	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	// time.Sleep(time.Second * 5)
	wg.Wait()
	fmt.Println("Main function finished")

}
