package main

import "fmt"

func main()  {
	var (
		a int = 10
		b int = 2
		addition int
		substraction int
		multiplication int
		division int
	)

	addition = a + b
	substraction = a - b 
	multiplication = a * b
	division = a / b

	fmt.Println(addition)
	fmt.Println(substraction)
	fmt.Println(multiplication)
	fmt.Println(division)
}