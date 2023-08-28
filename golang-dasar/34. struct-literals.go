/**
- struct adalah sebuah template data yang digunakan untuk menggabungkan
nol atau lebih tipe data lainnya dalam satu kesatuan
- struct biasanya representasi data dalam program aplikasi yang kita buat
- data di struct disimpan dalam field
- sederhananya struct adalah kumpulan dari field.package golangdasar
*/

/**
- struct adalah template data atau prototype data
- struct tidak bisa langsung digunakan
- namun kita bisa membuat data/object dari struct yang telah kita buat
*/

package main

import "fmt"

type Customer struct{
	Name, Address string
	Age			  int
}

func main() {
	// var wahyu Customer
	// wahyu.Name = "Wahyu Adi Kusuma"
	// wahyu.Address = "Yogyakarta, Indonesia"
	// wahyu.Age = 24

	// fmt.Println(wahyu.Name)
	// fmt.Println(wahyu.Address)
	// fmt.Println(wahyu.Age)

	var joko Customer = Customer{
		Name: "Joko",
		Address: "Solo, Indonesia",
		Age: 60,
	}

	fmt.Println(joko.Name)
	fmt.Println(joko.Address)
	fmt.Println(joko.Age)

	dimas := Customer {
		Name: "Dimas Adijaya",
		Address: "Jakarta, Indonesia",
		Age: 20,
	}

	fmt.Println(dimas.Name)
	fmt.Println(dimas.Address)
	fmt.Println(dimas.Age)

	adi := Customer{"Adijaya Pratama", "Yogyakarta", 15}
	fmt.Println(adi.Name)
	fmt.Println(adi.Address)
	fmt.Println(adi.Age)
}