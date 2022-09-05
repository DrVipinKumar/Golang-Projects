package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Output:")
	fmt.Println(strings.TrimRight("ABABBA", "BA"))
	fmt.Println(strings.TrimSuffix("ABABBA", "BA"))
}
