package main

import (
	"fmt"
	"runtime"
)

type Man struct {
	Name string
}

// Disable pass by value
// sehingga tidak membuat pointer baru tapi menggunakan pointer lama
func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	PrintMemUsage()
	angga := Man{"Angga"}
	angga.Married()

	fmt.Println(angga)
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
