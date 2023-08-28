package main

import "fmt"

func main() {
	for i := 2; i <= 100; i++ {
		if (i%2 == 1) {
			continue
		}
		// kalau ganjil maka operasi dibawah ini tidak dilakukan,
		// langsung ke next step
		fmt.Println(i)
	}
}