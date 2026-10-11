package main

import "fmt"

func main() {
	name := "irfan"
	var school string = "high school"
	var class = "5th semster bscs"

	fmt.Println(name, school, class)

	//lenght of strings
	size := len(name)
	fmt.Println(size, len(name), size == len(name))

	//comparring strings
	nikName := "Arfan"
	fmt.Println(len(nikName) == len(name))

	//accessing strings as index

}
