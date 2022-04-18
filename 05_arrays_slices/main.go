package main

import "fmt"

func main() {
	//Arrays
	var fruitArr [2]string

	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	foodArr := [2]string{"Apple", "banana"}

	fmt.Println(foodArr)

	//Slices

	anotherFruitArr := []string{"Apple", "guava", "mango", "cherry"}

	fmt.Println(anotherFruitArr)
	fmt.Println(len(anotherFruitArr))
	fmt.Println(anotherFruitArr[1:2])
}
