package main

import "fmt"

func main() {
	counter := 0

	increment := func() {
		// tetapi variabel counter diatas bisa diakses dari dalam function (2)
		counter++
		// variabel a hanya bisa diakses didalam function (1)
		a := "a"
		fmt.Println(a)
	}

	// fmt.Println(a)

	increment()

	fmt.Println(counter)
}
