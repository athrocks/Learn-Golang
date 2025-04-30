package main

import "fmt"

// Here’s a function that takes two ints and returns their sum as an int.
func add(a int, b int) int {
	return a + b
}

// When you have multiple consecutive parameters of the same type,
func plusPlus(a, b, c int) int {
	return a + b + c
}

// multiple return values
func getLang() (string, string) {
	return "Go", "Js"
}

// variadic functions
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func processIt(fn func(a int) int) {
	fmt.Println(fn(1))
}

func processIt2() func(a int) int {
	return func(a int) int {
		return 400
	}
}

func main() {
	fmt.Println("adding 1 + 2 =", add(1, 2))

	res := plusPlus(1, 2, 3)
	fmt.Println("adding 1 + 2 + 3 =", res)

	a, b := getLang()
	fmt.Println(a)
	fmt.Println(b)

	// multiple return values
	_, c := getLang()
	fmt.Println(c) // ignore the first return value
	// The blank identifier _ is used to ignore the first return value.

	// variadic functions
	sum(1, 2)
	sum(1, 2, 3)
	sum(1, 2, 3, 4)

	fn := func(a int) int {
		return 2
	}
	processIt(fn)

	res2 := processIt2()
	fmt.Println(res2(0)) // 400
}
