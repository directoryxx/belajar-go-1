package main

import "fmt"

func main()  {
	var (
		nama1 string
		nama2 string
		compare bool
	)

	nama1 = "Angga"
	nama2 = "Angga"

	compare = nama1 == nama2

	fmt.Println(compare)
}