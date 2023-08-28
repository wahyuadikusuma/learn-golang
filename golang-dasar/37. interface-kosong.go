package main

import "fmt"

func Benda() interface{} {
	return "halooo"
}
/**
func Benda(i int) interface{} { //bisa gini
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Apa benda"
	}
}
*/

func main() {
	// var wahyu interface{} = Benda()
	wahyu := Benda() // bisa gini
	fmt.Println(wahyu)
}