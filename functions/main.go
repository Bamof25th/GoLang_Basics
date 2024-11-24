package main

import "fmt"

func simpleFunc() {
	fmt.Println("Hey bro how are you now a days!")
}

func add(a, b int) (result int) {
	result = a + b
	return
}

func mult(a, b int) (result int) {
	result = a * b
	return
}

func main() {
	fmt.Println("we are learning GOLang")

	simpleFunc()
	ans := add(3, 4)
	mult := mult(3, 4)
	fmt.Println("add of two number is ", ans)
	fmt.Println("mult of two number is ", mult)
}
