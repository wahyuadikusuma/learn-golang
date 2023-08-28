/*
recover adalah function yang bisa kita gunakan
untuk mengangkap data panic
dengan recover proses panic akan terhenti, sehingga program akan tetap berjalan
*/
package main

import (
	"fmt"
)

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("Terjadi Error:", message)
	}
	fmt.Println("Aplikasi terhenti")
}

func runApp(error bool){
	defer endApp()
	if error {
		panic("Aplikasi ERROR")
	}
	fmt.Println("Aplikasi berjalan")
}

func main() {
	runApp(true)
	fmt.Println("Wahyu")
}