package main

import (
	"fmt"
	"strings"
)

func main()  {
	fmt.Println(strings.Contains("Angga Wijaya","Angga"))
	fmt.Println(strings.Contains("Angga Wijaya","Yoseps"))
	fmt.Println(strings.Split("Angga Wijaya"," "))
	fmt.Println(strings.ToLower("Angga Wijaya"))
	fmt.Println(strings.ToUpper("Angga Wijaya"))
	fmt.Println(strings.ToTitle("Angga Wijaya"))
	fmt.Println(strings.Trim("         Angga Wijaya "," "))
	fmt.Println(strings.ReplaceAll("Angga Wijaya","Angga","Test"))
}
