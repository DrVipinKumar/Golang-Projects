package main

import "fmt"

func BubbleSort(items []int) {
	for i := 0; i < len(items); i++ {
		for j := 0; j < len(items)-i-1; j++ {
			if items[j] > items[j+1] {
				temp := items[j]
				items[j] = items[j+1]
				items[j+1] = temp
			}
		}
	}
}
func main() {
	items := []int{12, 3, 114, 56, 5, 167, 34, 10, 123, 46}
	fmt.Println(items)
	BubbleSort(items)
	fmt.Println(items)
}
