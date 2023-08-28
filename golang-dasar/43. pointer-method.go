package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	wahyu := Man{"Wahyu"}
	wahyu.Married()

	fmt.Println(wahyu.Name)
}