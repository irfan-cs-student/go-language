package main

import "fmt"

func main() {

	name := "Irfan"
	matric := 20
	inter := 79
	entryTest := 40
	hasDocuments := false

	if matric >= 80 && inter >= 75 && entryTest >= 70 && hasDocuments {

		fmt.Println("congratulations", name, "!")
		fmt.Println("admission conformed")

	} else {
		if entryTest < 70 {
			fmt.Println("entry test number are low")
		}
		if matric < 80 {
			fmt.Println("matric num are low")
		}
		if inter < 75 {
			fmt.Println("inter num are low")
		}
		if !hasDocuments {
			fmt.Println("Document missing")
		}
	}

}
