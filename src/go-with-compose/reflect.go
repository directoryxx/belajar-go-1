package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	// structtag mirip dengan validation di laravel
	Name string `required:"true" max:"10"`
}

// Membuat validasi dengan structtag
func isValid(data interface{}) bool  {
	t := reflect.TypeOf(data)
	for i := 0; i< t.NumField();i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true"{
			value := reflect.ValueOf(data).Field(i).Interface()
			if value == ""{
				return false
			}
		}
	}

	return true
}

func main()  {
	angga := Sample{""}

	var sampleType reflect.Type = reflect.TypeOf(angga)

	fmt.Println(sampleType.NumField())
	fmt.Println(sampleType.Field(0).Name)
	fmt.Println(sampleType.Field(0).Tag.Get("required"))
	fmt.Println(sampleType.Field(0).Tag.Get("max"))

	fmt.Println(isValid(angga))
}
