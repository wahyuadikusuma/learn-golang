// yang sebelumnya di ignore karena materinya go-path
// sedangkan go-path sudah deprecated, digantikan go-mod

package main

import (
	"fmt"
	"strconv"
)

func main() {
	number, err := strconv.ParseInt("200", 10, 32)
	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}

	// pake Atoi : string -> int
	valueInt, _ := strconv.Atoi("100")
	if err == nil {
		fmt.Println(valueInt)
	} else {
		fmt.Println(err.Error())
	}

	// pake Itoa untuk parsing dari : int -> string
	valueString := strconv.Itoa(200) // returnnya cuman 1
	fmt.Println(valueString)
}