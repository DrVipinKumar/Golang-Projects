package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

func main() {
	file, err := os.OpenFile("diamonds.csv", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Panic(err)
	}
	df := dataframe.ReadCSV(file)
	fprice := df.Filter(dataframe.F{Colname: "price", Comparator: series.Greater, Comparando: 10000})
	fmt.Println(fprice)
	fcut := df.Filter(dataframe.F{Colname: "cut", Comparator: series.Eq, Comparando: "Premium"})
	fmt.Println(fcut)
}
