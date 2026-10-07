package main

import "fmt"

type Animal struct {
	name string
}
type Reboot struct {
	name string
}

//using methods
func (a Animal) run(speed int) {
	fmt.Println(a.name, "runs at", speed)
}
func (r Reboot) run(speed int) {
	fmt.Println(r.name, "run at speeds of ", speed)
}

//using interfaces
type Runner interface {
	run(speed int)
}

func startRunning(r Runner, speed int) {
	r.run(speed)
}

func main() {
	dog := Animal{"dog"}
	cat := Animal{"cat"} //struct

	animal := []Animal{ //slice
		{"kangro"},
		{"duk"},
		{"parrot"},
		{"sparrow"},
	}

	//changing type
	reboot := Reboot{"hollo reboot"}

	dog.run(12)
	startRunning(cat, 10)
	startRunning(dog, 15)
	startRunning(animal[0], 33)
	startRunning(reboot, 99)
}
