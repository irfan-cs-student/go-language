package main

import "fmt"

func main() {

	fmt.Println(`
	-------pattern 1-----
	`)

	for i := 1; i <= 5; i++ {

		for j := 1; j <= i; j++ {
			fmt.Print(j)
		}

		fmt.Println()
	}

	fmt.Println(`
	-------pattern 2-----
	`)

	for i := 1; i <= 5; i++ {

		for j := 1; j <= i; j++ {
			fmt.Print(j)
		}

		fmt.Println()
	}

}
