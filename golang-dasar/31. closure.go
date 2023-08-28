package main

import "fmt"

func main() {
	name := "Wahyu"
	counter := 0

	increment := func() {
		name = "Adi" // kalo gini bakal nge replace variable yg diatas
		// name := "Adi" // kalo gini dianggap variable yg beda dengan variable diatas
		fmt.Println("increment")
		counter++
	}


	increment()
	increment()
	fmt.Println(counter)
	fmt.Println(name)
}