package main

// doubly linked list

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()

	data.PushBack("Wahyu")
	data.PushBack("Adi")
	data.PushBack("Kusuma")

	fmt.Println(data.Front().Value)
	// fmt.Println(data.Front().Next().Value)
	// fmt.Println(data.Back().Value)

	// depan ke belakang
	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}

	// reverse - belakang ke depan
	for element := data.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
	}
}