package main

import "fmt"

func main() {
	var src, dst []int
	src = []int{1, 2, 3}
	//dst = make([]int, len(src))
	//copy(dst, src)
	dst = append(dst, src...)
	fmt.Println("dst:", dst)
}
