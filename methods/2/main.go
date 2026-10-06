package main

import "fmt"

type Class struct {
	name string
	age  int
}

func alter(class Class, name string) {

	class.name = name
	fmt.Println("from alter:", class)
	fmt.Println()

}
func alterName(class *Class, name string) {

	class.name = name
	fmt.Println("from alterName:", class)
	fmt.Println()

}

//by methods //parameter gets copy not adress
func (class Class) change(name string) {
	class.name = name
	fmt.Println("from change:", class)
	fmt.Println()

}

//by methods //parameter gets copy not adress
func (class *Class) changeName(name string) {
	class.name = name
	fmt.Println("from changeName:", class)
	fmt.Println()

}

func main() {

	class_1 := []Class{

		{"irfan", 22},
		{"shafi", 27},
		{"ali", 20},
	}
	fmt.Println("from main______:", class_1[0])
	fmt.Println()

	alter(class_1[0], "mahmood")

	fmt.Println("from main______:", class_1[0])
	fmt.Println()
	alterName(&class_1[0], "subhan")
	fmt.Println("from main:", class_1[0])
	fmt.Println()

	//by methods
	fmt.Println("______by methods chaning name____")
	class_1[0].change("noman")

	fmt.Println("now in main_______")
	fmt.Println(class_1)

	//by methods pointer
	fmt.Println("_________")
	class_1[0].changeName("loqman")

	fmt.Println("now in main_______")
	fmt.Println(class_1)

}
