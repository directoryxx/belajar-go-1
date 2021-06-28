package main

import (
	"fmt"
	"go-with-compose/database"

	// blank identifier berarti untuk mencegah package dihapus jadi biar tetap bisa digunakan
	//_ "go-with-compose/database"
)

func main()  {
	result := database.GetDatabase()
	fmt.Println(result)
}
