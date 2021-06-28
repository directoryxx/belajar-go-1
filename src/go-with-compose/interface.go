package main

import "fmt"

/*
	line 8 - 14 create interface
*/
type HasName interface {
	GetName() string
}

func sayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

/*
	line 19 - 25 implementasi interface
*/
type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

func main() {
	// line 29 - 32 exec interface
	var angga Person
	angga.Name = "Angga"

	sayHello(angga)
}
