package main

import "fmt"

type Product struct {
	name  string
	price int
}

func (x Product) Alt(name string) {
	x.name = name
	fmt.Println(x)
}
func (x *Product) change(name string) {

	x.name = name
	fmt.Println(x)
}

func alterName(abc *[]Product, name string) {
	(*abc)[0].name = name
	fmt.Println("printing from simple func _______")
	fmt.Println(abc)
	fmt.Println()
}

func main() {

	a := []Product{
		{"torch", 4},
		{"airpods", 70},
	}
	fmt.Println("printing ___main func________")
	fmt.Println(a)
	fmt.Println()

	alterName(&a, "gun")
	fmt.Println("printing ___main func________")
	fmt.Println(a)

	fmt.Println("methods_____introoo ")
	a[0].Alt("pods")

	fmt.Println("methods____as diferent techniques ")
	a[0].change("iphone")
	fmt.Println("In main", a)

}
