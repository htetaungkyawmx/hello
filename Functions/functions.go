package main

import "fmt"

func main() {
	x, y := 5, 6           // Declare two variables: x = 5, y = 6
	fmt.Println(add(x, y)) // Call the add function and print the result
}

func add(num1, num2 int) int {
	return num1 + num2 // Return the sum of num1 and num2
}
