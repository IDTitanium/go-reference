package main

import "fmt"

func main() {
	ids := []int{23, 234, 23, 244, 345, 34}

	// for i, id := range ids {
	// 	fmt.Printf("%d - ID: %d\n", i, id)
	// }

	// //Not using index
	// for _, id := range ids {
	// 	fmt.Printf("ID: %d\n", id)
	// }

	sum := 0
	for _, id := range ids {
		sum += id
	}

	fmt.Println("Sum", sum)

	//Range with maps
	anotherEmails := map[string]string{"bob": "bob@gmail.com", "sharon": "sharon@gmail.con"}

	for k, v := range anotherEmails {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range anotherEmails {
		fmt.Printf("%s\n", k)
	}
}
