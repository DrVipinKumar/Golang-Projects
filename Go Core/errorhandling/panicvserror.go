package main

import (
	"errors"
	"fmt"
)

func checkMySQL() bool {
	return false
}
func checkDBConnection() error {
	if checkMySQL() {
		return nil
	} else {
		return errors.New("MySQL is not running...")
	}
}
func main() {

	if err := checkDBConnection(); err != nil {
		//fmt.Println(err.Error())
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Panic Error:%s\n", r)
				fmt.Println("Backup server is running...")
				fmt.Println("My work in going on backup server...")
			}
		}()
		panic("MySQL is not running, run it first and start your program...")
	} else {
		fmt.Println("MySQL is running....")
	}
	fmt.Println("You can do your transaction in MySQL....")
}
