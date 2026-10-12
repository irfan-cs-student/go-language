package main

import "fmt"

func main() {
	//  defer executes arguments immediately

	x := 10

	defer fmt.Println(x)
	//It gets the arguments immediately, but executes the function later
	x = 20
	fmt.Println(x)

}
