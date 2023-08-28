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
	var wahyu Customer
	wahyu.Name = "Wahyu Adi Kusuma"
	wahyu.Address = "Yogyakarta, Indonesia"
	wahyu.Age = 24

	fmt.Println(wahyu.Name)
	fmt.Println(wahyu.Address)
	fmt.Println(wahyu.Age)
}