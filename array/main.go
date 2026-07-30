package main

import "fmt"

func main() {

	var num = [10]int{3, 2, 4, 4}
	name := [7]string{"irfan", "saeed", "ali"}
	fmt.Println("----With index------")
	fmt.Println(name[1])
	fmt.Println(num[2])

	fmt.Println("----Without index------")
	fmt.Println(name)
	fmt.Println(num)

	fmt.Println("---let practise-----")

	abc := [...]int{4, 4, 2, 4, 2, 45, 23}
	fmt.Println(abc)
	fmt.Println(abc[1], abc[1])
	fmt.Println("lenght of array", len(abc))
	fmt.Println("addin index --", abc[0]+abc[0])

	fmt.Println("---bool array-----")

	boolvalue := [...]bool{true, false, true}
	fmt.Println(boolvalue)
	fmt.Println(boolvalue[1])
	fmt.Println(boolvalue[1] == boolvalue[2])

}
