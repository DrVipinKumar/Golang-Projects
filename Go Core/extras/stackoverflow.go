//how to catch stackoverflow error in golang
package main

/*defer func() {
	if x := recover(); x != nil {
		log.Error("In recover, cought error====================", x)
	}
}()*/

//fn(xxx)
func RecoverPanic(msg string) {
	if r := recover(); r != nil {
		if msg != "" {
			if r != msg {
				panic(r)
			}
		}
	}
}

func foo() {
	defer RecoverPanic("stack overflow")

}
func main() {
	foo()
}
