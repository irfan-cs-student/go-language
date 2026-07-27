package main

import "fmt"

func greet() {
	fmt.Println("Wellcome to first coding functuion in go-lang")
}
func calcu(a, b int) (int, int) {
	return a + b, a - b
}
func introduce(name string, class string, age int) {
	fmt.Println("My name is :", name)
	fmt.Println("I am student of  :", class)
	fmt.Println("My age is :", age)
}
func square(n int) int {
	return n * n
}
func isEven(n int) {
	if n%2 == 0 {
		fmt.Println(n, ": Is even")
	} else {
		fmt.Println(n, ": Is odd")

	}

}
func even(n int) bool {
	if n%2 == 0 {
		return true
	} else {
		return false

	}
}
func isBigger(a, b int) {
	if a == b {
		fmt.Println("Both ", a, "and", b, "is equal")
		fmt.Println(a, "=", b)
	} else if a > b {
		fmt.Println(a, "is bigger than", b)
		fmt.Println(a, "=bigger")
		fmt.Println(b, "=smaller")

	} else {
		fmt.Println(a, "is smaller  than", b)
		fmt.Println(b, "=bigger")
		fmt.Println(a, "=smaller")

	}
	//calling another funtion
	//for chekin  whic is even or odd
	fmt.Println(a, "is even =", even(a))
	fmt.Println(b, "is even =", even(b))

}
func main() {

	//greet function
	greet()

	//addition function usinf 2 output return
	a, b := 2, 4
	sum, diff := calcu(a, b)
	fmt.Println("a+b :", sum)
	fmt.Println("a-b :", diff)

	//introduce
	introduce("irfan", "bS-CS", 22)

	//funtion returning sqaure of number
	n := 2
	fmt.Println("square of :", n, "=", square(n))

	//funtion  cheking for even or odd numbers
	num := 3
	isEven(num)          //no return
	cheking := even(num) //with bool return
	fmt.Println(num, "is even =", cheking)
	fmt.Println(num, "is even:", num%2 == 0)

	//funtion for cheking whisch is larger

	firstNum := 195
	secondNum := 42
	isBigger(firstNum, secondNum)
}
