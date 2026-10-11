package main

import "fmt"

func main() {
	greet := func(name string) {
		fmt.Println(name, "hello!")
	}
	greet("irfan")

}
