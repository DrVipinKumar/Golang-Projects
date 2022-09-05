package main

import (
	"fmt"
)

type dividebyzero struct {
	textError string
}

func (c dividebyzero) Error() string {
	return c.textError
}
func errorDivideByZero() error {
	return dividebyzero{"Divid by zero using Structure"}
}
func getDivision(num1, num2 int) (int, error) {
	if num2 == 0 {
		//return -1, errors.New("Divide by zero")
		return -1, errorDivideByZero()
	}
	return num1 / num2, nil
}
func main() {
	result, err := getDivision(20, 0)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("The division of two number is=", result)
	}
}
