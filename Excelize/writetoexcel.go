package main

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func main() {
	f := excelize.NewFile()
	//sheetname := f.NewSheet("Sheet2")
	f.SetCellValue("Sheet1", "A1", "Roll.No")
	f.SetCellValue("Sheet1", "B1", "Student Name")
	f.SetCellValue("Sheet1", "C1", "Course")
	f.SetCellValue("Sheet1", "A2", "1")
	f.SetCellValue("Sheet1", "B2", "Ravi Kumar")
	f.SetCellValue("Sheet1", "C2", "B.Tech")
	//	f.SetActiveSheet(sheetname)
	if err := f.SaveAs("Dr VipinKumar.xlsx"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("File Saved")
	}

}
