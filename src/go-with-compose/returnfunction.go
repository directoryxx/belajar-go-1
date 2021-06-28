package main

import "fmt"

// Parameter name adalah string dan pengembalian data harus didefinisikan
func getHello(name string) string {
	return "Hello " + name
}

func main() {
	result := getHello("Angga")

	fmt.Println(result)
}
