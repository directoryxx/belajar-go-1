package main

import "fmt"

func getGoodBye(name string) string {
	return "GoodBye " + name
}

func main() {
	// function value adalah keadaan yang memungkinkan untuk memasukkan function ke variabel
	// dan di set parameter dari variabel yang dibuat
	goodbye := getGoodBye
	fmt.Println(goodbye("angga"))
}
