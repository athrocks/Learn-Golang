package main

import (
	"fmt"
	"time"
)

func main() {

	// simple switch case
	i := 5

	// in go we do not need to write break statement in switch case
	// because it is automatically added at the end of each case
	switch i {
	case 1:
		fmt.Println("i is 1")
	case 2:
		fmt.Println("i is 2")
	case 3:
		fmt.Println("i is 3")
	default:
		fmt.Println("other")
	}

	// multiple condition switch

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	// type switch
	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case string:
			fmt.Println("I am a string:", t)
		case int:
			fmt.Println("I am an int:", t)
		case bool:
			fmt.Println("I am a bool:", t)
		default:
			fmt.Println("I don't know what I am:", t)
		}
	}

	whoAmI("hello")
	whoAmI(42)
	whoAmI(true)
	whoAmI(3.14)
}
