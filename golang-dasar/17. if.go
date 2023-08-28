package main

import "fmt"

func main() {
	name := "Kusuma"

	if name == "Wahyu" {
		fmt.Println("Halo Wahyu")
	} else {
		fmt.Println("Hi, Boleh kenalan?")
	}

	var length = len(name)
	if length > 5 {
		fmt.Println("Terlalu Panjang")
	} else {
		fmt.Println("Pas")
	}
}