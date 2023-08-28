package main

import "fmt"

func main() {
	var nilai32 int32 = 129
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32) 
	// var wahyu int8 = 128 //overflow -> error , karena maks int8 adalah 127

	// int8  rangenya	-128 ↔ 127
	// int16 rangenya	-32768 ↔ 32767
	// int32 rangenya	-2147483648 ↔ 2147483647

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8) // hasilnya akan menjadi -127, karena limitnya 127 dan kembali ke -128
	// fmt.Println(wahyu)

	//konversi string -> int -> string
	var (
		name = "Eko"
		e = name[0] //jadi uint 8
		eString = string(e)
	)

	fmt.Println()

	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)
}