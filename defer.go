package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication() {
	// defer berfungsi untuk memanggil suatu function di akhir function
	// meskipun ada error ditengah tengah
	defer logging()
	fmt.Println("Run Application")
	// fmt.Println(10 / 0)
}

func main() {
	runApplication()
}
