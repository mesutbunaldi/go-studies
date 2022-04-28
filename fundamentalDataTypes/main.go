package main

import "fmt"

func main() {

	var name = "Mesut"
	var age uint8 = 40
	var isMarried bool = false
	var weight = 72.2

	fmt.Println(name, age, isMarried, weight)

	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", isMarried)
	fmt.Printf("%T\n", weight)

}
