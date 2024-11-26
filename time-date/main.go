package main

import (
	"fmt"
	"time"
)

func main() {
	currTime := time.Now()

	fmt.Println("Current date and time:", currTime)
	fmt.Printf("Current date and time: %T\n", currTime)

	formatedData := currTime.Format("02-01-2006, Monday, 15:04 PM")

	fmt.Println("Formatted date:", formatedData)

	layoutString := "2006-01-02"
	dateString := "2026-11-21"

	formattedDate, _ := time.Parse(layoutString, dateString)

	fmt.Println("Parsed date:", formattedDate)

	new_date := currTime.Add(24 * time.Hour)

	fDate := new_date.Format("2006-01-02 Monday")

	fmt.Println("New date:", fDate)
}
