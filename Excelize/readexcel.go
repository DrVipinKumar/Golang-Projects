package main

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func main() {
	f, err := excelize.OpenFile("Dr VipinKumar.xlsx")
	if err != nil {
		fmt.Println(err)
	}
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
	}
	for _, row := range rows {
		for _, cell := range row {
			fmt.Print(cell, "\t")
		}
		fmt.Println()
	}

}
