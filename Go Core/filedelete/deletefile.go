package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	files, err := ioutil.ReadDir("files")
	if err != nil {
		fmt.Println(err)
	}
	for _, file := range files {
		match, _ := regexp.MatchString("json", file.Name())

		if match == true {
			filepath, err := filepath.Abs("./files/" + file.Name())
			if err != nil {
				fmt.Println(err)
			}
			err = os.Remove(filepath)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(file.Name() + " deleted")
			}

		}

	}

}
