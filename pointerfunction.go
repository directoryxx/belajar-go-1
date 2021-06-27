package main

import "fmt"

type Address struct {
	City, Province, Country string
}

// Mengubah data yang sebelumnya juga
// Tujuan jika struct memiliki field yang banyak
// bisa digunakan untuk menghemat memory
// karena setiap dipass golang akan membuat pointer baru
func changeProvince(address *Address, province string) {
	address.Province = province
}

func main() {
	address := Address{"Surabaya", "Jawa Timur", "Indonesia"}
	changeProvince(&address, "Jawa Barat")
	fmt.Println(address)
}
