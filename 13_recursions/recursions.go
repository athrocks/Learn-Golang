package main

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {

	println(fact(5)) // Output: 120
	println(fact(0)) // Output: 1
	println(fact(1)) // Output: 1
}
