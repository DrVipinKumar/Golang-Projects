package main

import "fmt"

func main() {
	var num [3]int
	num[0] = 1
	num[1] = 2
	num[2] = 3
	var snum []int = []int{1, 2, 4, 5}
	for i := 1; i < 5; i++ {
		snum = append(snum, i)

	}
	var newslice = num[2:3]
	var makeslice = make([]int, 3)
	makeslice[0] = 10
	makeslice[1] = 20
	makeslice[2] = 30
	makeslice = append(makeslice, 40)
	fmt.Println(makeslice)
	fmt.Println(newslice)
	fmt.Println(num)
	fmt.Println(snum)
	fmt.Printf("%T", num)
	fmt.Printf("\n%T", snum)

}
