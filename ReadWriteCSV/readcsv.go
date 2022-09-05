package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.OpenFile("students.csv", os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	fp := csv.NewReader(file)

	for {
		line, err := fp.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		fmt.Println(line)
	}
	defer file.Close()
}
