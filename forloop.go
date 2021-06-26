package main

import "fmt"

func main() {

	counter := 1

	for counter <= 10 {
		fmt.Println(counter)
		counter++
	}

	// Kayak Bahasa lain
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	slice := []string{"Nama", "aku", "angga"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	// For range untuk looping key value
	for index, name := range slice {
		fmt.Println("index ", index, " = ", name)
	}
}
