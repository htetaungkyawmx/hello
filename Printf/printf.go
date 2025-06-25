package main

import "fmt"

func main() {
	var name string = "Htet Aung Kyaw"
	const pi float64 = 3.1415926
	win := "test"
	x := 5

	fmt.Printf("%.3f \n", pi) // 3 decimal places of pi → 3.142
	fmt.Printf("%T \n", name) // Type of name → string
	fmt.Printf("%t \n", win)  // Boolean value → true
	fmt.Printf("%d \n", x)    // Integer → 5
	fmt.Printf("%b \n", 25)   // Binary of 25 → 11001
	fmt.Printf("%c \n", 34)   // Character of ASCII 34 → "
	fmt.Printf("%x \n", 15)   // Hexadecimal of 15 → f
	fmt.Printf("%e \n", pi)   // Scientific notation → 3.141593e+00

}
