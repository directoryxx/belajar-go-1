package main

import "fmt"

func factorial(value int) int {
	result := 1

	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorial(value-1)
	}
}

func main() {
	loop := factorial(5)
	fmt.Println(loop)

	recursiveLoop := factorialRecursive(5)
	fmt.Println(recursiveLoop)
}
