package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

// Binding method ke struct (1)
// Seakan akan bahwa function sayHi punya struct Customer
func (customer Customer) sayHi() {
	fmt.Println("Halo", customer.Name)
}

func main() {
	var angga Customer
	angga.Name = "Angga"
	angga.Address = "Sidoarjo"
	angga.Age = 23

	angga.sayHi()

}
