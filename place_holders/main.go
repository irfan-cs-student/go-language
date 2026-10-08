package main

import "fmt"

func main() {
	name := "Irfan"
	age := 22
	grad := 1
	class := 15

	fmt.Println(name, age, grad, class) //normal printing
	fmt.Printf("I'm %s . My age is %d. I made %d . my class is %d",
		name, age, grad, class)
}
