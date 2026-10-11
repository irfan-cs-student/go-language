package main

import "fmt"

func main() {

	greet := func(name string) {
		fmt.Println(name, "hello! ")
	}
	greet("irfan")

	func() {
		fmt.Println("________i love my country_______")
		results := 9
		fmt.Println(results + 9)
	}()
	//this imidately execute

	add := func(a, b int) int {
		return a + b
	}
	fmt.Println(add(2, 4))

	results := func(a, b int) int {
		return a * b
	}(2, 5)
	fmt.Println(results)

}
