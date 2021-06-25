package main

import "fmt"

func main() {
	var (
		names [4]string
	)

	names[0] = "nama"
	names[1] = "aku"
	names[2] = "adalah"
	names[3] = "angga"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
	fmt.Println(names[3])

	// Init array dengan isinya
	var values = [3]int{
		90,
		80,
		70,
	}

	fmt.Println(values)
}
