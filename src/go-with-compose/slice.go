package main

import (
	"fmt"
)

func main() {

	// ... berarti tidak tau berapa kapasitas array nya
	var month = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	// ambil array diambil dari index 4 sampai 6(sebelum 7)
	var slice1 = month[4:7]
	fmt.Println(slice1)
	// Panjang Slice
	fmt.Println(len(slice1))
	// Kapasitas slice dari index 4 - 12
	fmt.Println(cap(slice1))

	// Jika array diubah slice ikut berubah
	month[6] = "Diubah"
	fmt.Println(slice1)

	// Jika slice diubah array ikut berubah
	// Index 0 di slice = 4 di array month karena awal data slice langsung ambil index ke-4
	slice1[0] = "Coba"
	fmt.Println(slice1)

	slice2 := month[10:]
	fmt.Println(slice2)

	// Ketika menggunakan append berarti membuat array baru dari slice2
	// maka tidak mempengaruhi array yang berhubungan
	slice3 := append(slice2, "angga")
	fmt.Println(slice3)
	slice3[1] = "Bukan Desember"
	fmt.Println(slice3)
	fmt.Println(slice2)
	fmt.Println(month)

	// Membuat slice tanpa array
	newSlice := make([]string, 2, 5)
	newSlice[0] = "angga"
	newSlice[1] = "w"
	fmt.Println(newSlice)

	// copy slice
	// kapasitas harus sama supaya tidak kepotong
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	// Perbedaan Array dan slice
	// array harus didefinisikan panjangnya kalau slice tidak
	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
