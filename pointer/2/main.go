package main

import "fmt"

func change(a *int, b *int) {
	*a = *a + 5 //10+5=15----a
	*b = *a * 2 //20*2=30----b
	*a = *b - 3 //----30-3=27
}

func main() {
	x := 10
	y := 20

	change(&x, &y)

	fmt.Println(x) //27
	fmt.Println(y) //30
}
