package main

import (
	"fmt"
	"strconv"
)

// Define person struct

type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// Struct method (value receiver)

func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and i am " + strconv.Itoa(p.age)
}

//struct method (pointer receiver)

func (p *Person) hasBirthday() {
	p.age++
}

func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "f" {
		p.lastName = spouseLastName
	}
}

func main() {
	//Init person using struct

	person1 := Person{
		firstName: "Sama",
		lastName:  "Lasts",
		city:      "loas",
		gender:    "m",
		age:       12,
	}

	fmt.Println(person1)

	//Alternative

	person2 := Person{
		"Sama",
		"Lasts",
		"loas",
		"f",
		12,
	}

	fmt.Println(person2)

	fmt.Println(person1.firstName)
	person1.age = 10
	fmt.Println(person1)
	person1.hasBirthday()
	person2.getMarried("lawal")
	fmt.Println(person2.greet())
}
