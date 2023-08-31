package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile("w([a-z])u")

	fmt.Println(regex.MatchString("wahyu"))
	fmt.Println(regex.MatchString("wiu"))
	fmt.Println(regex.MatchString("wau"))

	search := regex.FindAllString("wao wiu wau wae",-1)
	fmt.Println(search)
}