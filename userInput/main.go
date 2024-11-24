package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("hey whats your name : ")
	// var name string
	// fmt.Scan(&name)
	// fmt.Println("Hello bruv :", name)

	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')

	fmt.Println("Hello bruv :", name)
}

// bufio is a package used to read line and take  standard inputs basically readline package in node js   