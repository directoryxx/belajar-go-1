package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}

	return total
}

func main() {
	fmt.Println(sumAll(1, 2, 3, 4, 5, 6, 7, 8))

	// insert slice di varargs
	slice := []int{1, 2, 3, 4, 535, 4, 3, 6}
	total := sumAll(slice...)
	fmt.Println(total)
}
