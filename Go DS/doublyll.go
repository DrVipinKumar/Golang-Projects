package main

import (
	"fmt"
	"os"
	"strconv"
)

type node struct {
	number int
	prev   *node
	next   *node
}
type doublylinkedlist struct {
	len  int
	head *node
	tail *node
}

func (dll *doublylinkedlist) InsertAtLast(num int) {
	var temp = &node{}
	temp.number = num
	temp.next = nil
	temp.prev = nil
	if dll.head == nil && dll.tail == nil {
		dll.head = temp
		dll.tail = temp
	} else {
		temp.prev = dll.tail
		dll.tail.next = temp
		dll.tail = temp
	}
	dll.len++
}
func (dll *doublylinkedlist) Display() {
	var p *node = dll.head
	for p != nil {
		fmt.Printf("%d ->", p.number)
		p = p.next
	}
}
func (dll *doublylinkedlist) DisplayReverse() {
	var p *node = dll.tail
	for p != nil {
		fmt.Printf("%d ->", p.number)
		p = p.prev
	}
}
func (dll *doublylinkedlist) InsertAtBegining(num int) {
	var temp = &node{}
	temp.number = num
	temp.prev = nil
	temp.next = nil
	if dll.head == nil && dll.tail == nil {
		dll.head = temp
		dll.tail = temp
	} else {
		temp.next = dll.head
		dll.head.prev = temp
		dll.head = temp
	}
	dll.len++
}
func (dll *doublylinkedlist) DeleteFromBegining() {
	if dll.head == dll.tail {
		dll.head = nil
		dll.tail = nil
	} else {
		dll.head = dll.head.next
		dll.head.prev = nil
	}
	dll.len--
}
func (dll *doublylinkedlist) DeleteFromEnd() {
	if dll.head == dll.tail {
		dll.head = nil
		dll.tail = nil
	} else {
		dll.tail = dll.tail.prev
		dll.tail.next = nil

	}
	dll.len--
}
func (dll *doublylinkedlist) DeleteFromSpecific(position int) {
	if dll.len >= position {
		var p *node = dll.head
		for i := 1; i < position-1; i++ {
			p = p.next
		}
		p.next = p.next.next
		p.prev = p
		dll.len--

	} else {
		fmt.Println("Specific Position not found")
	}
}
func (dll *doublylinkedlist) InsertAtSpecificPosition(num, position int) {
	var temp = &node{}
	temp.next = nil
	temp.prev = nil
	temp.number = num
	var p *node = dll.head
	for i := 1; i < position; i++ {
		p = p.next
	}
	temp.next = p.next
	p.next = temp
	temp.prev = p
	dll.len++
}
func main() {
	var dll = &doublylinkedlist{}
	var choice string
	for true {
		fmt.Println("Enter your choice")
		fmt.Println("1. Insert node at end")
		fmt.Println("2. Traverse/Display")
		fmt.Println("3. Traverse/Display in reverse order")
		fmt.Println("4. Insert at Begining")
		fmt.Println("5. Delete from Begining")
		fmt.Println("6. Delete from end")
		fmt.Println("7. Delete from specific position")
		fmt.Println("8. Insert at specific position")
		fmt.Println("9. Exit")
		fmt.Scanln(&choice)
		switch choice {
		case "1":
			var data string
			fmt.Println("Enter number")
			fmt.Scanln(&data)
			num, _ := strconv.Atoi(data)
			dll.InsertAtLast(num)
		case "2":
			dll.Display()
		case "3":
			dll.DisplayReverse()
		case "4":
			var data string
			fmt.Println("Enter number")
			fmt.Scanln(&data)
			num, _ := strconv.Atoi(data)
			dll.InsertAtBegining(num)
		case "5":
			dll.DeleteFromBegining()
		case "6":
			dll.DeleteFromEnd()
		case "7":
			var sposition string
			fmt.Println("Enter the position to delete")
			fmt.Scanln(&sposition)
			position, _ := strconv.Atoi(sposition)
			dll.DeleteFromSpecific(position)
		case "8":
			var data string
			var sposition string
			fmt.Println("Enter number")
			fmt.Scanln(&data)
			num, _ := strconv.Atoi(data)
			fmt.Println("Enter the position to delete")
			fmt.Scanln(&sposition)
			position, _ := strconv.Atoi(sposition)
			dll.InsertAtSpecificPosition(num, position)
		default:
			os.Exit(0)
		}
	}

}
