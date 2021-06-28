package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {

		// break menghentikan perulangan termasuk syntax dibawahnya
		if i == 5 {
			break
		}
		fmt.Println(i)
	}

	for i := 0; i < 10; i++ {

		// continue untuk skip iterasi
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}
