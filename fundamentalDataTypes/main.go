package main

import "fmt"

func main() {

	var name = "Mesut"
	var age int = 40
	var isMarried bool = false

	fmt.Println(name, age, isMarried)

	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", isMarried)

}
