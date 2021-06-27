package main

import (
	"errors"
	"fmt"
)

func Pembagi(value int, divide int) (int, error) {
	if divide == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	} else {
		result := value / divide
		return result, nil
	}
}

func main() {
	result, err := Pembagi(100, 0)

	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println("Error", err.Error())
	}
}
