package main

import (
	"fmt"
	"os"
)

func createFile() {
	f, err := os.Create("info.txt")
	defer f.Close()
	if err != nil {
		panic(err)
	}
	data := []byte("this is my information for text file")
	f.Write(data)
	f.WriteString("\nthis is another way to write information in text file")
}
func openFile() {
	f, err := os.Open("info.txt")
	if err != nil {
		panic(err)
	}
	data := make([]byte, 100)
	f.Read(data)
	fmt.Println(string(data))
}
func openFile2() {
	data, err := os.ReadFile("info.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
func main() {
	createFile()
	openFile()
	openFile2()
}
