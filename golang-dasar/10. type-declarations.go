package main

import "fmt"

func main() {
	type NoKTP string //membuat alias
	type Married bool

	var noKtpEko NoKTP = "128732149380001"
	var marriedStatus Married = true
	fmt.Println(noKtpEko, marriedStatus)
}