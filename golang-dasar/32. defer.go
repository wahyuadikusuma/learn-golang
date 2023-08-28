/*
defer function adalah func yg bisa kita jadwalkan
untuk dieksekusi setelah sebuah func selesai dieksekusi

defer func akan selalu dieksekusi walaupun terjadi error di func yang dieksekusi
*/
package main

import "fmt"

func logging(){
	fmt.Println("Selesai dieksekusi")
}

func runApplication(value int){
	// kalau defer di taruh di akhir, ketika terjadi error, 
	// dia tidak akan jalan. karena langsung exit (panic)
	defer logging() 
	fmt.Println("Run Application")
	result := 10 / value
	fmt.Println("Result:", result)
}

func main() {
	runApplication(0)
}