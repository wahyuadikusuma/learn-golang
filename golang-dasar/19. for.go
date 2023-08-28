package main

import "fmt"

func main() {
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("perulangan ke :", counter)
	}

	slice := []string{"Wahyu", "Adi", "Kusuma"}

	for i := 0; i<len(slice); i++{
		fmt.Println(slice[i])
	}

	for i, value := range slice {
		fmt.Println("index", i, "=", value)
	}

	var book map[string]string = make(map[string]string)
	book["title"] = "belajar Go-Lang"
	book["author"] = "Wahyu"
	book["kunci"] = "Salah"

	for key, value := range book {
		fmt.Println("key", key, "=", value)
	}
}