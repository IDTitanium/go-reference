package main

import "fmt"

func main() {
	//Define map
	emails := make(map[string]string)

	//Assign values
	emails["bob"] = "bob@gmail.co"
	emails["sharon"] = "hsaron@gmail.co"

	fmt.Println(emails["bob"])
	fmt.Println(len(emails))

	delete(emails, "bob")

	anotherEmails := map[string]string{"bob": "bob@gmail.com", "sharon": "sharon@gmail.con"}

	fmt.Println(anotherEmails)

	fmt.Println(emails)
}
