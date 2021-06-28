package main

import "fmt"


// mendefinisikan var apa saja yang di return
func getFullName() (firstName, middleName, lastName string) {
	firstName = "A.A"
	middleName = "Angga"
	lastName = "Wijaya"
	return
}

func main() {
	firstName, middleName, lastName := getFullName()

	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)
}
