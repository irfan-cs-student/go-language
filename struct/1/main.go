// 1. Student Struct

// Create a Student struct with:

// name → string
// age → int
// score → int

// Then:

// Create 2 students
// Give them different values
// Print their name, age, and score
// Change the score of one student
// Print the updated score

package main

import "fmt"

func main() {

	type student struct {
		name  string
		age   int
		score int
	}

	student_1 := student{
		name:  "irfan",
		age:   20,
		score: 80,
	}

	student_2 := student{
		name:  "usman",
		age:   19,
		score: 76,
	}

	fmt.Println("_______1st student______")
	fmt.Println("name: ", student_1.name)
	fmt.Println("age:", student_1.age)

	fmt.Println("score:", student_1.score)
	fmt.Println()
	fmt.Println("_______updating values of age ______")
	student_1.age = 22
	fmt.Println("age:", student_1.age)
	fmt.Println()

	fmt.Println("_______2nd student______")
	fmt.Println("name: ", student_2.name)
	fmt.Println("age:", student_2.age)
	fmt.Println("score:", student_2.score)
	fmt.Println()
	fmt.Println("_______updating values of score ______")
	student_2.score = 14
	fmt.Println("score:", student_2.score)
	fmt.Println()

}
