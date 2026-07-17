package main

import "fmt"

func main() {

	i := 1

	for i <= 9 {
		fmt.Println(i)
		i++

	}
	fmt.Println(i) //increment value

	// secon loop
	fmt.Println("-------x loop------")

	x := 1

	for x < 5 {

		fmt.Print(x, " ")
		x++
	}

	fmt.Println(x) //increment  value

	//2 variables in condition
	fmt.Println("3rd loop")

	for a, b := 1, 7; a < 7 && b > 1; {
		fmt.Println(a, " ", b)
		a++
		b++
	}
}

// go run loops/loops.go
