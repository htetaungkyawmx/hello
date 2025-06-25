package main

import "fmt"

func main() {

	//A constant is a variable  with a value that can't be changed
	const pi float64 = 3.1415926

	//You can declare multiple variables like this
	var (
		varA = 2
		varB = 3
	)

	fmt.Println(varA, varB)

	//String are a series of characters between "or"
	var Name string = "Htet Aung Kyaw"

	//Get string length
	fmt.Println(len(Name))

}
