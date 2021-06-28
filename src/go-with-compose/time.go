package main

import (
	"fmt"
	"time"
)

func main() {
	//fmt.Println(time.Now())
	//fmt.Println(time.Date(2020,8,20,1,1,1,0,time.UTC))

	layout := "2010-01-02"
	parse, _ := time.Parse(layout,"1993-08-10")
	fmt.Println(parse)
}
