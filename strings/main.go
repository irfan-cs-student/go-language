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
	fmt.Println(nikName[0])     //it's prints numeric value
	fmt.Printf("%c\n", name[0]) //it's prints charater

	// name[0] = 'A' // Compile error
	// ---strings are immutable means ,individaul charatr does'nt change

	//combing strings
	first_name, second_name := "muhammad", "irfan"

	full_name := first_name + "" + second_name

	fmt.Println(full_name)
	fmt.Println(first_name + "" + second_name)
	fmt.Println()

}
