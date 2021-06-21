package main

import "fmt"

func main() {
	var (
		nilaiakhir    int = 90
		presensi      int = 90
		lulus         bool
		luluspresensi bool
		lulusna       bool
	)

	luluspresensi = presensi > 80
	lulusna = nilaiakhir > 80

	lulus = luluspresensi && lulusna

	fmt.Println(lulus)

}
