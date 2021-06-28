package main

import "fmt"

func endApp() {
	// recover adalah untuk menangkap panic mirip try catch (2)
	// recover ditaruh di defer karena defer pasti dijalankan diakhir (2)
	// kalau recover dijalan dibawah panic maka recover tidak berjalan (2)
	message := recover()
	if message != nil {
		fmt.Println("Error message ", message)
	}
	fmt.Println("Aplikasi selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		// Panic lebih mirip ke error exception (1)
		// ketika panic maka semua yang dibawah panic tidak dijalankan (1)
		panic("Aplikasi Error")
	}
	fmt.Println("Aplikasi Berjalan")
}

func main() {
	runApp(true)
}
