package main

import (
	"fmt"
	"os"
	"strconv"
)

type stack struct {
	value int
	next  *stack
}

func (node *stack) push(num int) *stack {
	var temp = &stack{}
	temp.value = num
	temp.next = node
	node = temp
	return node

}
func (node *stack) pop() *stack {
	node = node.next
	if node == nil {
		fmt.Println("Stack is empty")
	}
	return node

}
func (node *stack) peek() {
	fmt.Println(node.value)
}
func (node *stack) display() {
	var p *stack = node
	for p.next != nil {
		fmt.Printf("| %d |", p.value)
		fmt.Println("\n_____")
		p = p.next
	}
}
func main() {
	head := new(stack)
	var choice string
	for true {
		fmt.Println("Enter your choice")
		fmt.Println("1. Push value in stack")
		fmt.Println("2. Pop value from stack")
		fmt.Println("3. Peep value from stack")
		fmt.Println("4. Display stack")
		fmt.Println("5. Exit")
		fmt.Scanln(&choice)
		switch choice {
		case "1":
			var data string
			fmt.Println("Enter value to push")
			fmt.Scanln(&data)
			num, _ := strconv.Atoi(data)
			head = head.push(num)
		case "2":
			head = head.pop()
		case "3":
			head.peek()
		case "4":
			head.display()
		default:
			os.Exit(0)
		}
	}
}
