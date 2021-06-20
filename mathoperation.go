package main

import "fmt"

func main() {
	var (
		a int = 10
		b int = 2
		// addition int
		// substraction int
		// multiplication int
		// division int
	)

	// Old way
	// addition = a + b
	// substraction = a - b
	// multiplication = a * b
	// division = a / b

	// New Way
	a += b
	fmt.Println(a)
	a -= b
	fmt.Println(a)
	a *= b
	fmt.Println(a)
	a /= b
	fmt.Println(a)

	// fmt.Println(addition)
	// fmt.Println(substraction)
	// fmt.Println(multiplication)
	// fmt.Println(division)
}
