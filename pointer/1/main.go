package main

import "fmt"

func main() {
	a := 3
	fmt.Println("a =", a)

	//pointer
	z := &a
	fmt.Println("&a =", &a)
	fmt.Println("z =", z)
	fmt.Println("&Z =", &z)
	fmt.Println("*Z =", *z)

	*z = -1
	fmt.Println("&a =", &a)
	fmt.Println("z =", z)
	fmt.Println("&Z =", &z)
	fmt.Println("*Z =", *z)
	fmt.Println("a =", a)

	b := &z
	fmt.Println("b =", b)
	fmt.Println("*b =", *b)

}
