package main

import "fmt"

func main() {
	student := make(map[string]int)

	student["s1"] = 32
	student["s2"] = 34

	fmt.Println(student)

	student2 := make(map[int]string)

	student2[1] = "s1"
	student2[2] = "s2"

	a1, b1 := student2[1]
	fmt.Println(a1, b1)

	for a2, b2 := range student2 {
		fmt.Println(a2, b2)
	}

}
