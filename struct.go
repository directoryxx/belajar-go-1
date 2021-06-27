package main

import "fmt"

// Struct mirip dengan class di OOP Language (1)
// Tujuan struct mendefinisikan blueprint terkait field (1)
// Tetapi struct tidak bisa langsung di gunakan (1)
type Customer struct {
	// Dibawah mirip dengan attribute (2)
	Name, Address string
	Age           int
	// Trigger Error cara ke-3
	// Married bool
}

func main() {
	// Struct diperlakukan seperti sebuah tipe data baru (3)
	var angga Customer
	// assign data ke blueprint struct (4)
	angga.Name = "Angga"
	angga.Address = "Sidoarjo"
	angga.Age = 23

	fmt.Println(angga.Name)

	// Alternatif assign struct (5)
	budi := Customer{
		Name:    "Budi",
		Address: "Surabaya",
		Age:     10,
	}

	fmt.Println(budi)

	// Alternatif assign struct (6)
	// Assign ini sangat rawan karena jika struct nya berubah (6)
	// Cara ini akan error (6)
	rudi := Customer{"Rudi", "Jember", 10}
	fmt.Println(rudi)
}
