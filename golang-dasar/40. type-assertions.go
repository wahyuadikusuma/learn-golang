package main

import "fmt"

func random() interface{} {
	return 64
}

// func main() {
// 	var result interface{} = random()
// 	var resultString string = result.(string)
// 	// jangan sampai salah ngelakuin type assertions,
// 	// kalau salah nanti error
// 	fmt.Println(resultString)
// }

// kalau mau lebih aman bisa pake switch
func main() {
	var result interface{} = random()
	switch value := result.(type) {
	case int:
		fmt.Println("Value", value, "is integer")
	
	case string:
		fmt.Println("Value", value, "is string")
	
	case bool:
		fmt.Println("Value", value, "is boolean")

	default:
		fmt.Println("Unknown type")
	}
}