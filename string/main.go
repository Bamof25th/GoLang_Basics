package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "hello, world, hey, aniket"

	arrStr := strings.Split(str, " ,")

	fmt.Println(arrStr)

	data := " two twe onr three two seven red green two yellow"

	numberStr := strings.Count(data, "two")

	fmt.Println(numberStr)

	stringStr := "        Hello, go!        "
	fstr := strings.TrimSpace(stringStr)

	fmt.Println(fstr)

	res := strings.Join([]string{str,data}, " ")
	fmt.Println(res)
}
