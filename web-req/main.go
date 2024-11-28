package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// learning a web req

	res, err := http.Get("https://dummyjson.com/users/1")
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer res.Body.Close()
	fmt.Println("Success", res.Status)

	buff, err := ioutil.ReadAll(res.Body)
	if err != nil {
			fmt.Println("Error" , err)
			return
	}
	fmt.Println("Response:", string(buff))

}
