package main

import (
	"fmt"
	"os"
	"strconv"
)

type node struct {
	number int
	next   *node
}
type queue struct {
	len  int
	font *node
	rear *node
}

func (head *queue) Enqueue(num int) {
	var temp = &node{}
	temp.number = num
	temp.next = nil
	if head.font == nil && head.rear == nil {
		head.font = temp
		head.rear = temp
	} else {
		head.rear.next = temp
		head.rear = temp
	}
	head.len++
}
func (head *queue) Dequeue() {
	if head.rear == head.font {
		head.rear = nil
		head.font = nil
	} else {
		head.font = head.font.next
	}
	head.len--
}
func (head *queue) Front() {
	fmt.Println("Front Node is =", head.font.number)
}
func (head *queue) Size() {
	fmt.Println("Size of Queue =", head.len)
}
func (head *queue) Display() {
	var p *node = head.font
	for p != nil {
		fmt.Printf("%d <-", p.number)
		p = p.next
	}
}
func main() {
	var myqueue = &queue{}
	var choice string
	for true {
		fmt.Println("Enter your choice")
		fmt.Println("1. Enqueue")
		fmt.Println("2. Dequeue")
		fmt.Println("3. Front")
		fmt.Println("4. Size")
		fmt.Println("5. Travese/Display")
		fmt.Println("6. Exit")
		fmt.Scanln(&choice)
		switch choice {
		case "1":
			var data string
			fmt.Println("Enter number for node in Queue")
			fmt.Scanln(&data)
			num, _ := strconv.Atoi(data)
			myqueue.Enqueue(num)
		case "2":
			myqueue.Dequeue()
		case "3":
			myqueue.Front()
		case "4":
			myqueue.Size()
		case "5":
			myqueue.Display()
		default:
			os.Exit(0)

		}
	}
}
