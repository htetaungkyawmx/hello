/*package main

import "fmt"

func main() {

	/*var EvenNum [5]int

	EvenNum[0] = 0
	EvenNum[1] = 2
	EvenNum[2] = 4
	EvenNum[3] = 6
	EvenNum[4] = 8

	fmt.Println(EvenNum[2])*/

// EvenNum := [5]int{0, 2, 4, 6, 8}

/*for i, value := range EvenNum {

	fmt.Println(value, i)
}

numSlice := []int{5, 4, 3, 2, 1}

sliced := numSlice[0:]
fmt.Println(sliced)

slice2 := make([]int, 5, 10)
copy(slice2, numSlice)
fmt.Println(numSlice)

slice3 := append(numSlice, 3, 0, -1)
fmt.Println(slice3)
*/

package main

import "fmt"

func main() {

	// Array: fixed size, holds 5 even numbers
	EvenNum := [5]int{0, 2, 4, 6, 8}

	// Loop through the array using range (index and value)
	for i, value := range EvenNum {
		fmt.Println(value, i) // Print value and index
	}

	// Slice: more flexible than array
	numSlice := []int{5, 4, 3, 2, 1}

	// Slice from index 0 to end (same as whole slice)
	sliced := numSlice[0:]
	fmt.Println(sliced) // Output: [5 4 3 2 1]

	// Create empty slice with length 5, capacity 10
	slice2 := make([]int, 5, 10)

	// Copy values from numSlice to slice2
	copy(slice2, numSlice)

	// Print original slice (numSlice)
	fmt.Println(numSlice) // Output: [5 4 3 2 1]

	// Add (append) 3 more values to numSlice
	slice3 := append(numSlice, 3, 0, -1)

	// Print new slice after append
	fmt.Println(slice3) // Output: [5 4 3 2 1 3 0 -1]
}
