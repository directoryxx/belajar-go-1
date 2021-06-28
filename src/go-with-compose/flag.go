package main

import (
	"flag"
	"fmt"
)

func main() {
	var host *string = flag.String("host","localhost","Put your database host")
	var username *string = flag.String("username","root","Put your database user")
	var password *string = flag.String("password","root","Put your database password")

	flag.Parse()

	// dikasih bintang karena host,username,password adalah pointer
	fmt.Println("Host :",*host)
	fmt.Println("Username :",*username)
	fmt.Println("Password :",*password)

	// ketika running cara run nya
	// go run flag.go -host=192.168.1.1 -username=test -password=test
}
