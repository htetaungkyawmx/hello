package main

import (
	"fmt"
	"io/ioutil" // Used to read files
	"log"       // Used for error logging
	"os"        // Used to create/write files
)

func main() {
	file, err := os.Create("sample.txt")

	if err != nil {
		log.Fatal(err) // Stop program if there's an error
	}

	file.WriteString("Hello My name is HAK and this file was created using GO!!")
	file.Close() // Always close the file after writing

	stream, err := ioutil.ReadFile("sample.txt")

	if err != nil {
		log.Fatal(err)
	}

	s1 := string(stream) // Convert bytes to string
	fmt.Println(s1)      // Print to console
}
