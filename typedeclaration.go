package main

import "fmt"


func main()  {
	// create alias tipe data
	type isMarried bool

	var marriedStatus isMarried = true

	fmt.Println(marriedStatus)
}