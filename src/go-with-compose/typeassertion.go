package main

import "fmt"

func random() interface{} {
	return "Error"
}

// Type Assertion mengubah interface kosong menjadi tipe data lain
// ex: string
// tetapi pastikan ketika mengkonversi harus sesuai dengan tipe datanya
// ketika tidak sama program akan crash
func main() {
	var result interface{} = random()
	var resultString string = result.(string)
	fmt.Println(resultString)

	// Safe way to check data type
	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Integer", value)
	default:
		fmt.Println("Unknown")
	}

}
