package main

import "fmt"

func main() {
	person := map[string]string{
		"name":   "Wahyu",
		"adress": "Solo",
	}
	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(len(person))
	
	var book map[string]string = make(map[string]string)
	book["title"] = "belajar Go-Lang"
	book["author"] = "Wahyu"
	book["kunci"] = "Salah"
	
	fmt.Println(book)
	delete(book, "kunci")
	fmt.Println(book)

}