package main

import "fmt"

// define a struct named Person

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Contact struct {
	Email string
	Phone float64
}

type Address struct {
	House int
	Area  string
	State string
}

type Employee struct {
	Person_Details Person
	Person_Contact Contact
	Person_Address Address
}

func main() {
	// var prince Person
	// fmt.Println("Prince Person :", prince)
	// prince.FirstName = "Prince"
	// prince.LastName = "Oberoy"
	// prince.Age = 23
	// fmt.Println("Details Person :", prince)

	person1 := Person{
		FirstName: "Aniket",
		LastName:  "Baghel",
		Age:       23,
	}

	fmt.Println("Person1 details :", person1)

	var employee1 Employee

	employee1.Person_Details = Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       35,
	}
	employee1.Person_Contact = Contact{
		Email: "johndoe@example.com",
		Phone: 9876543210,
	}
	employee1.Person_Address = Address{
		House: 123,
		Area:  "New York",
		State: "NY",
	}

	fmt.Println("Employee1 details :", employee1.Person_Address)
}
