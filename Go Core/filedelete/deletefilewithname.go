package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	files, err := ioutil.ReadDir("files")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		//fmt.Println(file.Name(), file.IsDir())
		match, _ := regexp.MatchString("json", file.Name())
		if match == true {
			dir, err := filepath.Abs("./files/" + file.Name())
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(dir)
			err = os.Remove(dir)
			if err != nil {
				fmt.Println(err)

			} else {
				fmt.Println("File successfully deleted")
			}

		}

	}
}
