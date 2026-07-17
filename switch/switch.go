package main

import "fmt"

func main() {

	num := 4
	num = 5

	switch num {
	case 1:
		fmt.Println("num 1")
	case 2:
		fmt.Println("num 2")
	case 3:
		fmt.Println("num 3")
	case 4:
		fmt.Println("num 4")
	case 5:
		fmt.Println("num 5")
	}

	marks := 80

	switch {

	case marks >= 90:
		fmt.Println("A+")

	case marks >= 80:
		fmt.Println("A")

	case marks >= 70:
		fmt.Println("B")

	default:
		fmt.Println("Fail")
	}

	var role string

	switch role {
	case "engineer":
		fmt.Println("he is enginner")
	case "admin":
		fmt.Println("he is admin")
	default:
		fmt.Println("he is random person")
	}
}

//  go run switch/switch.go
