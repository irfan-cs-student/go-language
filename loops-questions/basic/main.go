package main

import "fmt"

func main() {

	// Question#1 :Print numbers from 1 to 10.

	for i := 1; i <= 10; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	//----------------------------------------------
	// Question#2 :Print numbers from 10 to 1.

	for i := 10; i >= 1; i-- {
		fmt.Print(i, " ")
	}
	fmt.Println()

	//----------------------------------------------
	// Question#3 :Print only even numbers from 1 to 20.
	fmt.Println("Even numbers")

	for i := 0; i <= 20; i++ {

		if i%2 == 0 {
			fmt.Print(i, " ")
		}

		// 	OR
		// 	if i%2 != 0 {
		// 		continue
		// 	}
		// 	fmt.Print(i, " ")
		// }
	}
	fmt.Println()

	//----------------------------------------------
	// Question#4 :odd number
	fmt.Println("Odd numbers")

	for i := 0; i <= 20; i++ {

		if i%2 == 0 {
			continue
		}
		fmt.Print(i, " ")

	}
	fmt.Println()

	//----------------------------------------------
	// Question#5 :Find the sum of numbers from 1 to 100.
	fmt.Println("Find the sum of numbers from 1 to 100.")

	num := 0
	for i := 0; i <= 20; i++ {

		fmt.Println(num, "+", i, "=", num+i)
		num += i

	}
	fmt.Println("sum of 1 to 100 = ", num)

	fmt.Println()

	//----------------------------------------------
	// Question#6 :Find the factorial of a number.
	fmt.Println("Find the factorial of a number.")

	var n int
	fmt.Println("Enter a num")
	fmt.Scan(&n)

	factorial := 1

	for i := 1; i <= n; i++ {

		factorial *= i

	}
	fmt.Println("Factorial of ", n, "=", factorial)

	fmt.Println()

	//----------------------------------------------
	// Question#7 :Print the multiplication table of a number.
	fmt.Println("Print the multiplication table of a number.")

	var table_num int
	fmt.Println("Enter a num")
	fmt.Scan(&table_num)
	fmt.Println("Table of :", table_num)

	for i := 1; i <= 20; i++ {

		fmt.Println(table_num, "*", i, "=", table_num*i)

	}

	fmt.Println()

	//----------------------------------------------
	// Question#8 :Count digits in a number.
	fmt.Println("Count digits in a number.")

	var number int
	fmt.Print("Enter number: ")
	fmt.Scan(&number)
	count := 0

	if number == 0 {
		count = 1
	} else {
		for number != 0 {
			number /= 10
			count++
		}
	}
	fmt.Println("Digit in number: ", count)

	fmt.Println()

}
