package main

import "fmt"

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func main() {
	firstName := "Adi"
	sayHelloTo(firstName, "Kusuma")
	sayHelloTo("Wahyu", "Adi")
}