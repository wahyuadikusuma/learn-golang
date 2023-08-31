package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

// ini merupakan konsep circular linked list

func main() {
	var data *ring.Ring = ring.New(5)
	
	for i := 0; i < data.Len(); i++ {
		data.Value = "Data " + strconv.FormatInt(int64(i), 10)
		data = data.Next()
	}

	data.Do(func(value interface{})  {
		fmt.Println(value)
	})
	// jangan pakai for loop, karena ini circular linked list
	// selalu berputar, sehingga program tidak akan berhenti
}