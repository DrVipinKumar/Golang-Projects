package main

import (
	"fmt"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/xuri/excelize/v2"
)

func main() {
	f, err := excelize.OpenFile("Fruits.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	sheetname := f.NewSheet("Sheet2")
	f.SetActiveSheet(sheetname)
	// Insert a picture.
	if err := f.AddPicture("Sheet2", "A2", "image.jpg", ""); err != nil {
		fmt.Println(err)
	}
	// // Insert a picture to worksheet with scaling.
	if err := f.AddPicture("Sheet2", "J2", "image.jpg",
		`{"x_scale": 0.7, "y_scale": 0.7}`); err != nil {
		fmt.Println(err)
	}
	// // Insert a picture offset in the cell with printing support.
	if err := f.AddPicture("Sheet2", "Q2", "image.jpg", `{
	    "x_offset": 15,
	    "y_offset": 10,
	    "print_obj": true,
	    "lock_aspect_ratio": true,
	    "locked": false
	}`); err != nil {
		fmt.Println(err)
	}
	// Save the spreadsheet with the origin path.
	if err = f.Save(); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Image Inserted in Excel File")
}
