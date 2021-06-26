package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked")
	} else {
		fmt.Println("Welcome ", name)
	}
}

// Jika biasanya kita harus membuat function terlebih dahulu (1)
// Hal ini tidak diperlukan jika menggunakan anonymous function (1)
// func spamFilter(name string) string {
// 	if name == "Anjing" {
// 		return "..."
// 	} else {
// 		return name
// 	}
// }

func main() {
	// sebagai gantinya function dimasukkan ke dalam variabel (2)
	blacklist := func(name string) bool {
		return name == "Anjing"
	}

	registerUser("Angga", blacklist)

	// Alternatif (3)
	registerUser("Bgst", func(s string) bool { return s == "Bgst" })
}
