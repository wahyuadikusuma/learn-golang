package main

import "fmt"

func main() {

	// buat array secara manual
	var names [3]string

	names[0] = "Eko"
	names[1] = "Adi"
	names[2] = "Nugroho"
	// names[3] = "yu"  // error karena max size nya 3, sedangkan ini mengisi data ke-4

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
	// fmt.Println(names[3]) // error karena max size nya 3, sedangkan ini mengakses ke-4

	var values = [3]int8{
		90,
		80,
		75,
	}
	values[2] = 100
	fmt.Println(values)
	fmt.Println(values[1])

	fmt.Println(len(values))
}