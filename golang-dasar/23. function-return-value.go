package main

import "fmt"

func getHello(name string) string {
	if name == "" {
		return "Hello Bro!"
	} else {
		return "Hello " + name
	}
}

func main() {
	halo := getHello("Wahyu")
	fmt.Println(halo)

	fmt.Println(getHello("Kaizaki"))
}