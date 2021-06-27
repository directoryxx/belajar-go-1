package helper

import "fmt"

// Function tidak boleh sama jika dalam satu package
func SayHello(name string) {
	fmt.Println("Halo", name)
}
