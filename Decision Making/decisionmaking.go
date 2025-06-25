package main

import "fmt"

func main() {

	age := 18

	/*if age > 18 {
		fmt.Println("Yes, you can't make a decision")
	} else {
		fmt.Println("No, you can't make a decision")
	}*/

	switch age {

	case 16:
		fmt.Println("Prepare for college")
	case 17:
		fmt.Println("Don't run after girls")
	case 19:
		fmt.Println("Get yourself a job!")
	default:
		fmt.Println("Are you even alive?")

	}
}
