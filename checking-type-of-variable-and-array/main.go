package main

import "fmt"

func main() {
	a := 32
	b := 'a'
	c := "irfan cs student of 4th semester"

	d := true
	num := [...]int{34, 55, 23, 10}

	fmt.Print("type of ", a, " = ")
	fmt.Printf("%T\n", a)

	fmt.Print("type of ", b, " = ")
	fmt.Printf("%T\n", b)

	fmt.Print("type of ", c, " = ")
	fmt.Printf("%T\n", c)

	fmt.Print("type of ", d, " = ")
	fmt.Printf("%T\n", d)

	fmt.Print("type of ", num, " = ")
	fmt.Printf("%T\n", num)

}
