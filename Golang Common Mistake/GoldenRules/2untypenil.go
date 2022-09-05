//use of untyped nil
package main

import "fmt"

func main() {
	//var x =nil
	var x interface{} = nil //error

	fmt.Println(x)
}
