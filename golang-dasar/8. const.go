package main

import "fmt"

func main() {
	const prefix string = "xocp"
	const name = "latihanwahyu"

	fmt.Println(prefix + "_" + name)

	const (
		firstName = "Wahyu"
		middleName = "Adi"
		occupation = "IT"
	)

	fmt.Println(firstName, middleName, occupation)
}