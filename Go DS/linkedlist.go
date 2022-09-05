package main

import (
	"fmt"
	"os"
	"strconv"
)

type linkelist struct {
	number int
	next   *linkelist
}

func (node *linkelist) insert(num int) {
	var temp = &linkelist{}
	temp.number = num
	temp.next = nil
	if node == nil {
		node = temp
	} else {
		var p = &linkelist{}
		p = node
		for p.next != nil {
			p = p.next
		}
		p.next = temp

	}

}
func (node *linkelist) display() {
	var p = &linkelist{}
	p = node.next
	for p != nil {
		fmt.Printf("%d ->", p.number)
		p = p.next
	}
}
func (node *linkelist) deletefirst() *linkelist {
	node = node.next
	return node
}
func (node *linkelist) deletelast() {
	var p *linkelist = node
	for p.next.next != nil {
		p = p.next
	}
	p.next = nil
}
func (node *linkelist) insertspecific(num, position int) {
	var p *linkelist = node
	var temp = &linkelist{}
	for i := 1; i < position; i++ {
		p = p.next
	}
	temp.number = num
	temp.next = p.next
	p.next = temp
}
func (node *linkelist) deletespecific(position int) {
	var p *linkelist = node
	for i := 1; i < position; i++ {
		p = p.next
	}
	p.next = p.next.next
}
func (node *linkelist) insertatbegining(num int) *linkelist {
	var temp = &linkelist{}
	temp.number = num
	temp.next = node.next
	node.next = temp
	return node
}
func main() {
	head := new(linkelist)
	var choice string
	for true {
		fmt.Println("\nEnter your choice")
		fmt.Println("1. Insert at last")
		fmt.Println("2. Display linked list")
		fmt.Println("3. Deleting from begining")
		fmt.Println("4. Deleting from last")
		fmt.Println("5. Deleting from specific position")
		fmt.Println("6. Insert at specific position")
		fmt.Println("7. Insert at begining")
		fmt.Println("3. Exit")
		fmt.Scanln(&choice)
		switch choice {
		case "1":
			var data string
			fmt.Println("Enter your value for linked list node")
			fmt.Scanln(&data)
			num, _ := strconv.Atoi(data)
			head.insert(num)
		case "2":
			head.display()
		case "3":
			head = head.deletefirst()
		case "4":
			head.deletelast()
		case "5":
			var sposition string
			fmt.Println("Enter position to delete")
			fmt.Scanln(&sposition)
			position, _ := strconv.Atoi(sposition)
			head.deletespecific(position)
		case "6":
			var sposition string
			var data string
			fmt.Println("Enter position to insert")
			fmt.Scanln(&sposition)
			position, _ := strconv.Atoi(sposition)
			fmt.Println("Enter value to add")
			fmt.Scanln(&data)
			num, _ := strconv.Atoi(data)
			head.insertspecific(num, position)
		case "7":
			var data string
			fmt.Println("Enter value to add")
			fmt.Scanln(&data)
			num, _ := strconv.Atoi(data)
			head = head.insertatbegining(num)
		default:
			os.Exit(0)
		}
	}

}
