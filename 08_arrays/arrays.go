package main

import "fmt"

func main() {
	var nums [5]int
	fmt.Println(nums)

	// array length
	fmt.Println(len(nums))

	nums[0] = 1
	fmt.Println(nums[0])

	var vals [3]bool
	fmt.Println(vals)

	var str [4]string
	fmt.Println(str)

	a := [3]int{11, 22, 33}
	fmt.Println(a)

	// You can also have the compiler count the number of elements for you with ...
	// The ... operator tells the compiler to count the number of elements in the array.
	b := [...]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	// If you specify the index with :, the elements in between will be zeroed.
	c := [...]int{100, 3: 400, 500}
	fmt.Println(c) // [100 0 0 400 500]

	// 2D array
	d := [2][2]int{
		{1, 2},
		{3, 4},
	}
	fmt.Println(d) // [[1 2] [3 4]]
}
