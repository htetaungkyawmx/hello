package main

import "fmt"

func main() {

	var name string = "Htet Aung Kyaw"
	const pi float64 = 3.1415926
	win := true
	x := 5

	/*fmt.Println(len(name))
	fmt.Println(name + " is a children")*/

	fmt.Println("%.3f \n", pi)
	fmt.Println("%T \n", name)
	fmt.Printf("%t \n", win)
	fmt.Printf("%d \n", x)
	fmt.Printf("%b \n", 25)
	fmt.Printf("%c \n", 34)
	fmt.Printf("%x \n", 15)
	fmt.Printf("%e \n", pi)

}
