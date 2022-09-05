package main

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func main() {
	f, err := excelize.OpenFile("MyStudent.xlsx")
	if err != nil {
		fmt.Println(err)
	}
	rows, err := f.GetRows("Sheet2")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Excel File Information are:")
	for _, row := range rows {
		for _, col := range row {
			fmt.Print(col, "\t")
		}
		fmt.Println()
	}

}
