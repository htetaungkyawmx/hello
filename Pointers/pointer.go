package main

import "fmt"

func main() {

	x := 10
	changeValue(&x)
	fmt.Println(x)

	y := 20
	method1(y)
	fmt.Println(y)

}

func changeValue(x *int) {
	*x = 7
}

func method1(y int) {
	y = 10
	fmt.Println("inside method1", y)
}
