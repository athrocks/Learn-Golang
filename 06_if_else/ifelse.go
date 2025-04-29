package main

import "fmt"

func main() {
	age := 19

	if age >= 18 && age < 22 {
		fmt.Println("You are an Teen.")
	} else if age >= 22 {
		fmt.Println("You are an Adult.")
	} else {
		fmt.Println("You are a Kid.")
	}

	var role = "admin"
	var hasPermission = true
	if role == "admin" && hasPermission {
		fmt.Println("You have access to the admin panel.")
	}

	// var declared in if statement
	if age := 19; age >= 18 {
		fmt.Println("You are an Adult of age", age)
	} else {
		fmt.Println("You are a Kid.")
	}

	// go does not have ternary operator like other languages

}
