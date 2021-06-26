package main

import "fmt"

// Membuat type declaration (alias) supaya jika terlalu panjang bisa langsung panggil alias
type Filter func(string) string

// function as parameter adalah memanggil function
// untuk dimasukkan sebagai parameter di function lain
func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello" + nameFiltered)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Angga", spamFilter)
}
