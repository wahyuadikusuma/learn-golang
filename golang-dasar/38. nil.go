package main

import "fmt"

func NewMap(name string, address string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
			"address": address,
		}
	}
}

func main() {
	person := NewMap("Wahyu", "Yogyakarta")
	// var person map[string]string = NewMap("Wahyu", "Yogyakarta") // bisa gini juga

	if person == nil {
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(person)
	}
}