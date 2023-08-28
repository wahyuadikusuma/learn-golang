package main

import "fmt"

func main() {
	var name1 = "Eko"
	var name2 = "Budi"

	fmt.Println(name1>name2)
	var result bool = name1 > name2 //keyword bool tidak wajib asal langsung diinisialisasi inputannya
	fmt.Println(result)
}