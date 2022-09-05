package main

import "fmt"

func SelectionSort(items []int) {
	for i := 0; i < len(items)-1; i++ {
		max := i
		for j := i + 1; j < len(items); j++ {
			if items[max] < items[j] {
				max = j
			}
		}
		temp := items[i]
		items[i] = items[max]
		items[max] = temp
	}
}
func main() {
	items := []int{212, 23, 34, 45, 35, 56, 38, 57, 90, 123, 67}
	fmt.Println(items)
	SelectionSort(items)
	fmt.Println(items)
}
