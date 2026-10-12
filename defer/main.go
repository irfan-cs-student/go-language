package main

// In Go, defer delays a function call until the surrounding function is about to finish.
import "fmt"

func main() {

	defer fmt.Println("first ")
	fmt.Println("second ")
	fmt.Println("third ")

}
