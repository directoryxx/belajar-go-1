package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// address1 := Address{"Sidoarjo", "Jawa Timur", "Indonesia"}
	// address2 := address1

	// address2.City = "Surabaya"

	// // Pada dasar golang itu pass by value
	// // Jadi value dari memory address1 dicopy ke memory address2
	// fmt.Println(address1) // Tidak berubah
	// fmt.Println(address2)

	// // Cara pass by reference
	// // Maksudnya adalah menggunakan pointer ke memory yang dicopy
	// address1 := Address{"Sidoarjo", "Jawa Timur", "Indonesia"}
	// address2 := &address1

	// address2.City = "Surabaya"

	// fmt.Println(address1) // berubah
	// fmt.Println(address2)

	// Force use new pointer
	// Maksudnya adalah memaksa pointer yang untuk menuju * pointer
	address1 := Address{"Sidoarjo", "Jawa Timur", "Indonesia"}
	address2 := &address1

	address2.City = "Surabaya"

	*address2 = Address{"Malang", "Jawa TImur", "Indonesia"}

	fmt.Println(address1) // berubah
	fmt.Println(address2)

	// Membuat pointer baru tetapi dengan struct kosong
	var address4 *Address = new(Address)
	address4.City = "Jogja"
	fmt.Println(address4)
}
