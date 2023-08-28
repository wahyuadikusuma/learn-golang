package main

import "fmt"

func endApp() {
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
}