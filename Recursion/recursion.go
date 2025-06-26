package main

import "fmt"

func main() {

	fmt.Println(factorial(5)) // Calls the factorial function with 5

}

func factorial(num int) int {

	if num == 0 { // Base case: 0! = 1
		return 1
	}
	return num * factorial(num-1) // Recursive call
}
