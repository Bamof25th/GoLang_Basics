package main

import (
	"fmt"
	"os"
)

// this helper allow our code to be more readable
func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	// file, error := os.Create("f1.txt")
	// if error != nil {
	// 	fmt.Println(error)
	// 	return
	// }
	// defer file.Close()

	// content := "hello world"
	// byte , err := io.WriteString(file, content + "\n")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("byte written: ", byte)
	// fmt.Println("successfully created a file")

	// file, err := os.Open("f1.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer file.Close()

	// //  create a buffer to read  the file content

	// buff := make([]byte, 1024)

	// // Read the file content
	// for {
	// 	n, err := file.Read(buff)

	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	fmt.Println(string(buff[:n]))
	// }

	content, err := os.ReadFile("f1.txt")
	check(err)
	fmt.Println(string(content))

}
