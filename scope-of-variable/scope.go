package main

import "fmt"

func main() {

	for a := 1; a < 8; a++ {

		fmt.Print(a, " ")

	}

	a := 1
	fmt.Println(a)
	{
		a = 2
		fmt.Println(a)

		a := a + 3
		fmt.Println(a)
	}
	fmt.Println(a)

	fmt.Println(`
	----------creating new variable--------`)

	b := 22
	fmt.Println(b)
	c := b

	{
		b := 22 + 22
		c = b
		fmt.Println(b)
		fmt.Println(c)

	}
	fmt.Println(c + 3)

	fmt.Println(b)
	fmt.Println("new challenge")

	// {
	// 	abc := 123
	// }
	// fmt.Println(abc)

}
