package main

import "fmt"

const age = 30

func main() {
	const name = "golang"

	// name = "js" // cannot assign to name

	const pi = 3.14

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)

	// grouping constants
	const (
		port = 8080
		host = "localhost"
	)

	fmt.Println("Port:", port, "Host:", host)

}
