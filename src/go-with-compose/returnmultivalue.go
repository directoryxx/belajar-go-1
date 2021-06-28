package main

import "fmt"

func multiValue() (string, string) {
	return "Angga", "W"
}

func main() {
	// inisiasi variabel langsung dua karena return variabel juga dua
	namaDepan, namaBelakang := multiValue()

	fmt.Println(namaDepan, namaBelakang)

	// ignore return kedua
	namaDepan1, _ := multiValue()

	fmt.Println(namaDepan1)
}
