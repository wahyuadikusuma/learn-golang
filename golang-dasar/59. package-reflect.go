package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Nama string `required:"true" max:"10"`
	Alamat string `required:"true" max:"100"`
}

func isValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			value := reflect.ValueOf(data).Field(i).Interface()
			if value == "" {
				return false
			}
		}
	}
	return true
}

func main() {
	sample := Sample{"Wahyu", "Yogya"}

	var sampleType reflect.Type = reflect.TypeOf(sample)

	fmt.Println(sampleType.NumField())
	fmt.Println(sampleType.Field(0).Name)
	structField := sampleType.Field(0) 
	required := structField.Tag.Get("required")
	fmt.Println(sampleType.Field(1).Tag.Get("max"))

	fmt.Println(required)
	fmt.Println(isValid(sample)) // true karena nama dan alamat bukan string kosong
}