package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}

func main() {
	total := sumAll(10, 90, 80, 30)
	fmt.Println(total)
	
	// jika ingin menggunakan parameter slice
	number := []int{1,2,3,4,5}
	total = sumAll(number...)
	fmt.Println(total)
}