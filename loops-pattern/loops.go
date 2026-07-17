package main

import "fmt"

func main() {

	fmt.Println(`
	-------pattern 1: Number Triangle-----
	`)

	for i := 1; i <= 5; i++ {

		for j := 1; j <= i; j++ {
			fmt.Print(j)
		}

		fmt.Println()
	}

	fmt.Println(`
	-------pattern 2:Same number-----
	`)

	for i := 1; i <= 5; i++ {

		for j := 1; j <= i; j++ {
			fmt.Print(i)
		}

		fmt.Println()
	}

	fmt.Println(`
	-------pattern 3:Same Diferent number-----
	`)

	num := 1
	for i := 1; i <= 5; i++ {

		for j := 1; j <= i; j++ {
			fmt.Print(num, " ")
			num++
		}

		fmt.Println()
	}

	fmt.Println(`
	-------pattern 4:Inverted Numbers-----
	`)

	for i := 5; i >= 1; i-- {

		for j := 1; j <= i; j++ {
			fmt.Print(j)

		}

		fmt.Println()
	}
	fmt.Println(`
	-------pattern 5:Pyramid-----
	`)

	n := 8
	for i := 1; i <= n; i++ {

		for j := 1; j <= n-i; j++ {
			fmt.Print(" ")

		}
		for s := 1; s <= 2*i-1; s++ {
			fmt.Print("*")

		}

		fmt.Println()
	}

	fmt.Println(`
	-------pattern 5:Diamond-----
	`)

	n = 8

	//uper dioamond part
	for i := 1; i <= n; i++ {

		for sp := 1; sp <= n-i; sp++ {
			fmt.Print(" ")

		}
		for s := 1; s <= 2*i-1; s++ {
			fmt.Print("*")

		}

		fmt.Println()
	}

	//lower diomond part
	for i := n - 1; i >= 1; i-- {

		for sp := 1; sp <= n-i; sp++ {
			fmt.Print(" ")

		}
		for s := 1; s <= 2*i-1; s++ {
			fmt.Print("*")

		}

		fmt.Println()
	}

}
