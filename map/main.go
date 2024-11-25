package main

import "fmt"

func main() {
	// map very similar to hash map
	// most used data structure in golang after slice
	studentGrades := make(map[string]int)

	studentGrades["Prince"] = 34
	studentGrades["Aniket"] = 39
	studentGrades["Ram"] = 34
	studentGrades["John"] = 35
	studentGrades["Geo"] = 38

	fmt.Println("hey heres the map  you made :", studentGrades)
	fmt.Println("Marks of Aniket is :", studentGrades["Aniket"])
	fmt.Println("Length of map : ", len(studentGrades))

	studentGrades["Aniket"] = 40
	fmt.Println("Marks of Aniket is :", studentGrades["Aniket"])

	delete(studentGrades, "Geo")
	fmt.Println("Marks of Geo is :", studentGrades["Geo"])

	// `	// clear(studentGrades)
	// 	// fmt.Println("Map is :", studentGrades)`

	value, exists := studentGrades["Ben"]

	fmt.Println("Davis is there or not:", exists)
	fmt.Println("Davis is there then his marks are:", value)

	marks, prs := studentGrades["Aniket"]

	fmt.Println("Aniket is there or not:", prs)
	fmt.Println("Aniket is there then his marks are:", marks)

	data := map[string]int{"man": 10, "women": 34, "unknown": 4}

	for name, val := range data {
		fmt.Printf("the gender is %s , the amount is %T\n", name, val)
	}

}
