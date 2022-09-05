package main

import (
	"fmt"
	"os"
)

func checkPath() error {
	var err *os.PathError = nil
	fmt.Println(err)
	return err
}
func main() {
	err := checkPath()
	fmt.Println(err)
	fmt.Println(err == (*os.PathError)(nil))
}
