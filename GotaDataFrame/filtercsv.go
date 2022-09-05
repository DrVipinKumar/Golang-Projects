package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-gota/gota/series"

	"github.com/go-gota/gota/dataframe"
)

func main() {
	file, err := os.OpenFile("diamonds.csv", os.O_RDWR, os.ModePerm)
	if err != nil {
		log.Panic(err)
	}
	df := dataframe.ReadCSV(file)
	//fmt.Println(df)
	fprice := df.Filter(dataframe.F{Colname: "price", Comparator: series.Greater, Comparando: 5000})
	fmt.Println(fprice)
	fcut := df.Filter(dataframe.F{Colname: "cut", Comparator: series.Eq, Comparando: "Ideal"})
	fmt.Println(fcut)
}
