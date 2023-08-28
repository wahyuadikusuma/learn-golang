package main

import "fmt"

func main() {
	var name string

	name = "Wahyu Adi"
	fmt.Println(name)
	
	name = "Wahyu Adi Kusuma"
	fmt.Println(name)

	var friendName = "rilo"
	fmt.Println(friendName)

	friendName = "100" // error kalau direplace integer
	fmt.Println(friendName)

	age := 31 // := menggantikan kata kunci var
	fmt.Println(age)

	var (
		satu = 1
		dua = 1
	)
	fmt.Println("satu + dua : ", satu + dua)
}