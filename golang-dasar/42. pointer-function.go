package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address){
	address.Country = "Indonesia"
}

func main() {
	// address1 := Address{"Sleman", "Yogyakarta", "Indonesia"}
	// address2 := address1 // pass by value (address1 dicopy ke address2)
	// address2 := &address1 // pass by reference (address1 dan address2 merefer data yang sama di memory yang sama)
	// address3 := &address1
	/**
	// deklarasi dengan cara lain:
	var address1 Address = Address{"Sleman", "Yogyakarta", "Indonesia"}
	var address2 *Address = &address1
	var address3 *Address = &address1
	*/

	var alamat = Address{
		City: "Wonogiri",
		Province: "Jawa Tengah",
		Country: "",
	}
	var alamatPointer *Address = &alamat
	fmt.Println("Sebelum diubah:",alamat)
	ChangeCountryToIndonesia(alamatPointer)
	ChangeCountryToIndonesia(&alamat) // sama aja kalau mau langsung
	fmt.Println("Setelah diubah:",alamatPointer)
	fmt.Println("Setelah diubah:",alamat)
}