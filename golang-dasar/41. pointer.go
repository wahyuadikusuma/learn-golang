package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Sleman", "Yogyakarta", "Indonesia"}
	// address2 := address1 // pass by value (address1 dicopy ke address2)
	address2 := &address1 // pass by reference (address1 dan address2 merefer data yang sama di memory yang sama)
	address3 := &address1
	/**
	// deklarasi dengan cara lain:
	var address1 Address = Address{"Sleman", "Yogyakarta", "Indonesia"}
	var address2 *Address = &address1
	var address3 *Address = &address1
	*/

	address2.City = "Bantul"

	// TANPA OPERATOR *
	// membuat address baru yang berbeda dengan address1
	// address2 = &Address{"Kulon Progo", "Yogyakarta", "Indonesia"} 
	
	// fmt.Println(address1)
	// fmt.Println(address2)
	// fmt.Println(address3)

	fmt.Println()

	// OPERATOR *
	// memaksa data yang merefer address1 pindah ke data baru 
	*address2 = Address{"Kulon Progo", "Yogyakarta", "Indonesia"} 

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	fmt.Println()

	// Kalau ingin membuat pointer dengan data kosong - dengan func new
	var address4 *Address = new(Address)
	address4.City = "Jakarta"
	fmt.Println(address4)

}