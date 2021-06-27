package helper

import "fmt"

// Access modifier kayak private,public,protected

var version = "1.0.0" // tidak bisa diexport karena huruf kecil
var Application = "golang"

// Function tidak boleh sama jika dalam satu package
func SayHello(name string) {
	fmt.Println("Halo", name)
}

// Tidak bisa diexport karena huruf kecil
func sayGoodBye(name string) {
	fmt.Println("Bye", name)
}
