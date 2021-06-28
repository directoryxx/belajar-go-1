package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println("Argument : ",args)
	hostname, err := os.Hostname()
	if err == nil{
		fmt.Println(hostname)
	}

	// Ambil system env
	// Test system env
	// export username=test
	// export password=test
	username := os.Getenv("username")
	password := os.Getenv("password")
	fmt.Println("Username",username)
	fmt.Println("Password",password)
}
