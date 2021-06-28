package main

import "fmt"

func main() {
	// Map menggunakan format key-value
	// map[string]string = tipe data string untuk key dan tipe data string untuk valuenya
	orang := map[string]string{
		"nama":   "angga",
		"alamat": "pinggir jalan",
	}

	// akses map sama kayak array
	fmt.Println(orang)
	fmt.Println(orang["nama"])
	fmt.Println(orang["alamat"])

	// tambah key
	orang["pekerjaan"] = "programmer"
	fmt.Println(orang)

	// panjang map
	fmt.Println(len(orang))

	// membuat map dengan make
	buku := make(map[string]string)
	buku["title"] = "Buku kosong"
	buku["pengarang"] = "ngarang"
	buku["salah"] = "salah"

	fmt.Println(buku)

	// hapus key
	delete(buku, "salah")
	fmt.Println(buku)
}
