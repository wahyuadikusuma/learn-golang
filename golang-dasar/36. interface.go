package main

import "fmt"

type HasName interface {
	GetName() string
}

func sayHello(hasName HasName) {
	fmt.Println("Hello,", hasName.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string{
	return person.Name
}

// try another example

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	var wahyu Person
	wahyu.Name = "Wahyu"
	sayHello(wahyu)

	cat := Animal {
		Name: "Pusss",
	}

	sayHello(cat)
}