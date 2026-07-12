package main

import "fmt"

func main() {

	price := 600
	isMember := true
	coupen := true

	if price >= 1000 {
		fmt.Println("Orignal price=700")
		fmt.Println("discount 20%")

	} else if price >= 500 && isMember {
		fmt.Println("Orignal price=700")
		fmt.Println("discount 10%")

	} else if coupen {
		fmt.Println("Orignal price=700")
		fmt.Println("discount 5%")
	} else {
		fmt.Println("Orignal price=700")
		fmt.Println("Discount 0%")
	}
}
