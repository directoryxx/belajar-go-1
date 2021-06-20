package main

import "fmt"

func main() {
	// init var
	var (
		nilai32 int32 = 100000000
		nilai64 int64
		nilai8  int8
	)
	// Do some magic prevent overflow increase value
	nilai64 = int64(nilai32) * 100

	fmt.Println(nilai64)

	// Decrease value to fit int8
	nilai8 = int8(nilai32) / 100

	fmt.Println(nilai8)

}
