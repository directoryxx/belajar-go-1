package main

import "fmt"

func main() {
	var name, uang string

	name = "Angga"
	uang = "banyak"

	fmt.Println(name)
	fmt.Println(uang)

	name = "Angga W"
	fmt.Println(name)

	// Langsung Init var tanpa tipe data
	var namaku = "angga"
	fmt.Println(namaku)

	// Ini variabel tanpa var
	namakuini := "angga"
	fmt.Println(namakuini)

	//Define variabel langsung banyak tidak assign value
	var (
		firstName string
		lastName  string
	)

	firstName = "Angga"
	lastName = "Wijaya"

	fmt.Println(firstName)
	fmt.Println(lastName)

	//Define variabel langsung banyak dengan assign value
	var (
		firstName1 = "Angga"
		lastName1  = "Wijaya"
	)

	fmt.Println(firstName1)
	fmt.Println(lastName1)
}
