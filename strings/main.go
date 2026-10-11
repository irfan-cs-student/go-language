package main

import (
	"fmt"
	"strings"
)

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

	//capital letters
	fmt.Println(strings.ToUpper(name))
	fmt.Println(strings.ToLower(name))
	fmt.Println(strings.ReplaceAll(name, "irfan", "Fann"))
	fmt.Println(strings.Contains(name, "irfan"))
	fmt.Println(strings.Contains(name, "Fann"))

	//string split
	sentence := "I love my country:"
	slice := strings.Split(sentence, " ") //it's now slice

	fmt.Println(slice)
	fmt.Println(slice[0])
	fmt.Println(slice[1])
	fmt.Println(slice[2])
	fmt.Println(slice[3])

	//joing strings
	words := []string{"I", "love", "Go"}

	result := strings.Join(words, "-")

	fmt.Println(result)
}
