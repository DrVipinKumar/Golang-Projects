package main

import (
	"fmt"
	"regexp"
)

func main() {
	r, _ := regexp.Compile("([0-9]*)")
	fmt.Println(r.MatchString("8888888888 hello 8987676543 india mobile number"))
	fmt.Println(r.FindAllString("8888888888 hello 8987676543 india", -1))
}
