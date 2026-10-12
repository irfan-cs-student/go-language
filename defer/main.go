package main

// In Go, defer delays a function call
//  until the surrounding function is about to finish.

import "fmt"

func main() {

	defer fmt.Println("first ")
	fmt.Println("second ")
	fmt.Println("third ")

	//defer with anonymous func
	defer func() {
		fmt.Println("defer with anonymous func")
	}()

	// Deferred calls execute in LIFO order (Last In, First Out).

	defer fmt.Println("A")
	fmt.Println("B")
	defer fmt.Println("C")

}
