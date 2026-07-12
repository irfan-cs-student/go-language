package main

import "fmt"

func main() {

	//print
	fmt.Print("hello ")
	fmt.Print("World ")

	//with print-line
	fmt.Println("hello")
	fmt.Println("world")

	//with comma
	fmt.Println("hello", "world")

	//diferent arguments
	fmt.Println("My name is:", "irfan", "\n My age is:", 22, "\n I am currently doing :", "BSCS")

	//Doing small math
	fmt.Println("ali takes apples: 5kg", "\n he takes food:40 rupess", "\n he spend 44 rupesse on dates", "\n total: ", "5kg * 100 rupesse =", 5*100, "+44 datres", "\n total : ", "500 + 44 =", 500+44)

	//backticks
	fmt.Println(`abc def "jhi" lmn 
	123
	456
	7879
	xyzzzzzzzzzzzzzzzzzz,zzzzzz`)
}
