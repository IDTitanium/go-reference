package main

import "fmt"

func main() {
	// var name string = "brad"
	var age int64 = 37
	const isCool = true

	//Shorthand declaration
	name := "brad"
	size := 1.3

	fmt.Println(name, age, isCool, size)
	fmt.Printf("%T\n", size)
}
