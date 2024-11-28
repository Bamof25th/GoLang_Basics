package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("learning url")

	myurl := "https://www.example.com/page?param1=value1&param2=value2"

	fmt.Printf("type of url is %T\n", myurl)

	parsedURL, err := url.Parse(myurl)
	if err != nil {
		fmt.Println("error parsing url", err)
		return
	}

	fmt.Printf("Type of my-url: %T\n", parsedURL)
	fmt.Println("protocol:", parsedURL.Scheme)
	fmt.Println("protocol:", parsedURL.Host)
	fmt.Println("protocol:", parsedURL.RawQuery)
	fmt.Println("protocol:", parsedURL.Path)
	fmt.Println("protocol:", parsedURL)

}
