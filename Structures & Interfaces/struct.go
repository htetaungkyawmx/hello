package main

import "fmt"

// 🔷 Interface - defines behavior (like a contract)
type Animal interface {
	speak() string
}

// 🔶 Struct - like a class with fields
type Dog struct {
	name string
}

// 🔶 Struct - another animal
type Cat struct {
	name string
}

// ✅ Dog implements Animal (by defining speak())
func (d Dog) speak() string {
	return d.name + " says Woof!"
}

// ✅ Cat implements Animal (by defining speak())
func (c Cat) speak() string {
	return c.name + " says Meow!"
}

// 🔁 A function that accepts any Animal
func makeItSpeak(a Animal) {
	fmt.Println(a.speak())
}

// 🚀 main function
func main() {
	dog := Dog{name: "Buddy"}
	cat := Cat{name: "Mimi"}

	// Use interface function
	makeItSpeak(dog)
	makeItSpeak(cat)
}
