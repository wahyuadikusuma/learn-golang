package main

import "fmt"

func getFullName() (string, string, string) {
	return "Bane", "Beban", "Banget" 
}

func main() {
	_, firstName, lastName := getFullName() // _ untuk meng ignore return value
	fmt.Println(firstName)
	fmt.Println(lastName)
}