package main

import "fmt"

// cara biasa

// func sayHelloWithFilter(name string, filter func(string) string) {
// 	nameFiltered := filter(name)
// 	fmt.Println("Hello", nameFiltered)
// }

// pake type declaration

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if (name == "anjing") || (name == "bege")  {
		return "..."
	} else {
		return name
	}
}

func main() { 
	sayHelloWithFilter("anjing", spamFilter)
	sayHelloWithFilter("Andi", spamFilter)
	sayHelloWithFilter("bege", spamFilter)
}