package main

import "fmt"

func main() {
	a := []int{10, 20, 30, 40, 50, 60, 70, 80}
	b := a[2:6]

	fmt.Println("lenght:", len(a), "\n", "capacity:", cap(a))

	fmt.Println(b)
	fmt.Println()

	//all clear about include exclude---

	c := b[1:3]
	fmt.Println(c)

	c = append(c, 100, 200)
	fmt.Println("after appending", c)
	fmt.Println("now b is :", b)
	fmt.Println("now a is :", a)

	fmt.Println("now b0 is 99")
	b[0] = 99
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Println("lenght:", len(a), "\n", "capacity:", cap(a))

}
