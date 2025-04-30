package main

import (
	"fmt"
	"slices"
)

// slices are dynamic arrays
// most used data structure in go
// slices are more powerful than arrays
func main() {

	// unintialized slice is nil
	var nums []int

	fmt.Println(nums)        // []
	fmt.Println(nums == nil) // true
	fmt.Println(len(nums))   // 0

	var nums2 = make([]int, 5)
	fmt.Println(nums2)        // [0 0 0 0 0]
	fmt.Println(nums2 == nil) // false
	fmt.Println(cap(nums2))

	// make([]T, len, cap) creates a slice of type T with length len and capacity cap
	// if cap is not provided, it will be equal to len
	var nums3 = make([]int, 3, 5)
	fmt.Println(nums3)        // [0 0 0]
	fmt.Println(nums3 == nil) // false
	fmt.Println(cap(nums3))   // 5

	nums3 = append(nums3, 1)
	fmt.Println(nums3)
	fmt.Println(cap(nums3))

	nums3 = append(nums3, 2, 3, 4, 5, 6)
	fmt.Println(nums3)      // [0 0 0 1 2 3 4 5 6]
	fmt.Println(cap(nums3)) // 10

	nums4 := []int{1, 2, 3, 4, 5}
	fmt.Println(nums4) // [1 2 3 4 5]
	fmt.Println(cap(nums4))
	fmt.Println(len(nums4))

	s := []string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(s) // [a b c d e f]

	l := s[1:4]    // [b c d]
	fmt.Println(l) // [b c d]

	m := s[:3]     // [a b c]
	fmt.Println(m) // [a b c]

	n := s[3:]     // [d e f]
	fmt.Println(n) // [d e f]

	c := make([]string, len(s))
	copy(c, s)     // copy s to c
	fmt.Println(c) // [a b c d e f]

	// slice package
	var nums5 = []int{1, 2}
	var nums6 = []int{1, 2}

	// fmt.Println(nums5 == nums6) // false, slices are not comparable
	fmt.Println(slices.Equal(nums5, nums6)) // true, use slices package to compare slices

	var num7 = [][]int{{1, 2}, {3, 4}}
	fmt.Println(num7) // [[1 2] [3 4]]

}
