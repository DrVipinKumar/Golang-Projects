package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	demo()

}

func demo() {
	// getting .txt file content
	txtByte, err := ioutil.ReadFile("./package.txt")
	if err != nil {
		log.Fatal(err)
	}
	txtString := string(txtByte)

	// getting main lines to append
	txtLines := getMainLines(txtString)

	// getting .json file content
	jsonByte, err := ioutil.ReadFile("./package.json")
	if err != nil {
		log.Fatal(err)
	}
	jsonString := string(jsonByte)

	// getting index of where the main lines will be append on .json file
	indexOfAppend := getIndex(jsonString, `"scripts": {`)

	// merging
	finalString := merge(jsonString, indexOfAppend, txtLines)

	fmt.Println(finalString)
	err = ioutil.WriteFile("./package.json", []byte(finalString), 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("successfully merged")
}

func getMainLines(txtString string) string {
	// eleminating `"scripts": {` part from .txt file
	txtString = strings.TrimSpace(strings.TrimPrefix(txtString, `"scripts": {`))

	// eleminating `}` part from .txt file
	txtString = strings.TrimSpace(strings.TrimSuffix(txtString, `}`))

	// finally return main lines
	return txtString
}

func getIndex(jsonString, subStr string) int {
	// identifying index of where to append
	index := strings.Index(jsonString, subStr)
	if index == -1 {
		fmt.Println("invalid sub-string")
	}

	// index is the value of first char of subStr
	// so adding the length upto `{`
	return index + 12 //  `"scripts": {` -->  have 12 chars
}

func merge(jsonString string, indexOfAppend int, txtLines string) string {
	// newStr is consist of 3 part
	// first: content of .json file upto `"scripts": {`
	// second: content of main lines
	// thirs: last content of .json file
	newStr := jsonString[:indexOfAppend] + txtLines + jsonString[indexOfAppend:]

	return newStr
}
