package main

import "fmt"

func main() {
	// Init Const dengan tipe data
	const firstName string = "Angga"
	// Init const tanpa tipe data
	const lastName = "Wijaya"

	fmt.Println(firstName)
	fmt.Println(lastName)

	// Line 11-12 akan error jika program running comment/hapus
	// firstName = "Angaa"
	// lastName = "Wijayaa"
}
