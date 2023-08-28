/**
- struct adalah tipe data seperti lainnya, dia bisa digunakan
sebagai parameter untuk function
- namun jika kita ingin menambahkan method ke dalam structs,
sehingga seakan-akan sebuah struct memiliki function
*/

package main

import "fmt"

type Customer struct{
	Name, Address string
	Age			  int
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hello,", name, "My Name is", customer.Name)
}

func main() {
	wahyu := Customer{Name: "Wahyu"}
	wahyu.sayHi("Brodi!")
}