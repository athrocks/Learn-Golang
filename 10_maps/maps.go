package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]int)

	m["a"] = 11
	m["b"] = 22
	m["c"] = 33

	fmt.Println("map:", m)
	fmt.Println("len:", len(m))

	fmt.Println("a:", m["a"])

	// if key does not exist, it will return 0 for int
	fmt.Println("h:", m["h"])

	// check if key exists
	v, ok := m["a"] // v has the value of m["a"], ok is true if key exists
	if ok {
		fmt.Println("a:", v)
	} else {
		fmt.Println("a not found")
	}

	fmt.Println("keys:")
	for k := range m {
		fmt.Println(k, m[k])
	}

	// delete a key
	delete(m, "a")
	fmt.Println("map:", m)

	// clear the map
	clear(m)
	fmt.Println("map:", m)

	mpp := map[string]int{
		"k1": 11,
		"k2": 22,
		"k3": 33,
	}
	fmt.Println("map:", mpp)

	m1 := map[string]int{"price": 100, "phones": 3}
	m2 := map[string]int{"price": 100, "phones": 3}

	// fmt.Println("m1 == m2:", m1 == m2) // maps are not comparable, so this will not work
	fmt.Println(maps.Equal(m1, m2)) // use maps package to compare maps
}
