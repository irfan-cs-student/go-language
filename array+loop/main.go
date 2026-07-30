package main

import "fmt"

func main() {
	num := [...]int{0: 1, 4: 23, 7}
	fmt.Println(num)

	name := [...]string{
		0: "irfan",
		3: "ali",
		4: "nasir",
		7: "rizwan",
		9: "rizwan",
	}
	fmt.Println(name)
	fmt.Println("lenght of name-array:", len(name))
	fmt.Print("type of name:")
	fmt.Printf("%T\n", name)

	if name[7] == name[9] {
		fmt.Print("same name")
	} else {
		fmt.Print("not same name")

	}

	fmt.Println(`
	-------------------learning array with loops-------`)

	for i := 0; i <= len(name)-1; i++ {
		fmt.Println("index:", i, "value:", name[i])
	}

	fmt.Println("-----------range loop------------")
	for _, value := range name {
		fmt.Println("value:", value)

	}
}
