package main

import (
	"fmt"
	"runtime"
)

type Address struct {
	City, Province, Country string
}

// Mengubah data yang sebelumnya juga
// Tujuan jika struct memiliki field yang banyak
// bisa digunakan untuk menghemat memory
// karena setiap dipass golang akan membuat pointer baru
func changeProvince(address *Address, province string) {
	address.Province = province
}

func main() {
	PrintMemUsage()
	address := Address{"Surabaya", "Jawa Timur", "Indonesia"}
	changeProvince(&address, "Jawa Barat")
	fmt.Println(address)
	PrintMemUsage()
}

// PrintMemUsage outputs the current, total and OS memory being used. As well as the number
// of garage collection cycles completed.
func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
