package main

import "fmt"

func main() {
	var a int = 10
	// you have to use any declared variable in the code, otherwise it will give an error
	fmt.Println(a)

	var name string = "golang"
	fmt.Println(name)

	// if type not given it will be inferred from the value
	var name2 = "golang also"
	fmt.Println(name2)

	// shorthand syntax for variable declaration
	name3 := "golang also"
	fmt.Println(name3)
	// shorthand syntax for variable declaration with type inference
}
