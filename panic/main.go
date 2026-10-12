// panic stops the normal execution of the current function
// and starts unwinding the call stack.
//
//	Deferred functions still execute while the stack unwinds.

package main

import "fmt"

func test() {
	fmt.Println("A")
	panic("Something went wrong")
	fmt.Println("B")
}

func main() {
	fmt.Println("Start")
	test()
	fmt.Println("End")
}
