package main

import "fmt"

//before going to interfaces practicing methods
type Animal struct {
	name string
}

//making custom type as slice
type janver []Animal

//using methods
func (janvar Animal) run(speed int) {

	fmt.Println(janvar.name, " move at speed of ", speed)
}

//method for custom type
func (animal janver) runs(speed int) {

	for i := range animal {

		fmt.Println(animal[i].name, "runs at speed of : ", speed)
		speed += 10
	}
}
func main() {

	dog := Animal{"dog"} //struct
	cat := Animal{"cat"} //struct

	animal := []Animal{ //slice
		{"kangro"},
		{"duk"},
		{"parrot"},
		{"sparrow"},
	}

	fmt.Println()
	fmt.Println("_________run method________")

	cat.run(19)
	dog.run(24)
	animal[0].run(25)

	//if upload whols slice
	janver(animal).runs(10)
}
