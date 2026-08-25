package main

import "fmt"

func main() {

	a := make([]int, 3, 5)
	fmt.Println("lenght of slice:", len(a))
	fmt.Println("Capacity of slice:", cap(a))

	a = append(a, 2)
	a = append(a, 22)
	a = append(a, 23)

	fmt.Println("lenght of slice:", len(a))
	fmt.Println("Capacity of slice:", cap(a))
	fmt.Println("slice values:", a)
	fmt.Println("slice values from 2 index to 5 index:", a[2:5])

	a = append(a, 30)
	fmt.Println(a)
	fmt.Println("lenght of slice:", len(a))
	fmt.Println("Capacity of slice:", cap(a))
	a = append(a, 30)
	a = append(a, 31)
	fmt.Println("lenght of slice:", len(a))
	fmt.Println("Capacity of slice:", cap(a))

}
