package main

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func main() {
	f := excelize.NewFile()
	f.SetCellValue("Sheet1", "A1", "Roll. No")
	f.SetCellValue("Sheet1", "B1", "Student Name")
	f.SetCellValue("Sheet1", "C1", "Course")
	f.SetCellValue("Sheet1", "A2", "1")
	f.SetCellValue("Sheet1", "B2", "Ravi Kumar")
	f.SetCellValue("Sheet1", "C2", "B.Tech")
	f.SetCellValue("Sheet1", "A3", "2")
	f.SetCellValue("Sheet1", "B3", "Sachin Kumar")
	f.SetCellValue("Sheet1", "C3", "M.Tech")
	sheet2 := f.NewSheet("Sheet2")
	f.SetCellValue("Sheet2", "A1", "Roll. No")
	f.SetCellValue("Sheet2", "B1", "Student Name")
	f.SetCellValue("Sheet2", "C1", "Course")
	f.SetCellValue("Sheet2", "A2", "1")
	f.SetCellValue("Sheet2", "B2", "Ravi Kumar")
	f.SetCellValue("Sheet2", "C2", "B.Tech")
	f.SetCellValue("Sheet2", "A3", "2")
	f.SetCellValue("Sheet2", "B3", "Sachin Kumar")
	f.SetCellValue("Sheet2", "C3", "M.Tech")
	f.SetActiveSheet(sheet2)
	err := f.SaveAs("MyStudent.xlsx")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("File Created and Data Inserted in MS Excel File")
	}

}
