package main

import "fmt"

// loop only has one syntax in Go, which is the same as the for loop in C/C++/Java.
func main() {
	// while loop
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// infinite loop
	// for {
	// 	fmt.Println("Hello, world!")
	// }

	// for loop
	for i := 1; i <= 3; i++ {
		fmt.Println(i)
	}

	// for loop with range
	for i, v := range []int{10, 21, 32} {
		fmt.Println(i, v)
	}

	for i := range 3 {
		fmt.Println(i)
	}

	// for loop with break
	for i := 1; i <= 3; i++ {
		if i == 2 {
			break
		}
		fmt.Println(i)
	}

	// for loop with continue
	for i := 1; i <= 3; i++ {
		if i == 2 {
			continue
		}
		fmt.Println(i)
	}

}
