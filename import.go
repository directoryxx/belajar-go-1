package main

import (
	"fmt"
	"go-with-compose/helper"
)

func main() {
	helper.SayHello("angga")
	fmt.Println(helper.Application)
	// fmt.Println(helper.version) // Error karena diawali huruf kecil
}
