package main

import (
	"fmt"
	"os"
)

// In Go, file handling means
// creating files, opening them, writing data,
// reading data, appending data, and deleting files.

func main() {

	file, _ := os.Create("myinfo.txt")
	defer file.Close()

	file.WriteString("my name is irfan")
	fmt.Println("file created")

	data, _ := os.ReadFile("myinfo.txt")

	fmt.Println(string(data))
	file.WriteString("I'm muslim")
	data2, _ := os.ReadFile("myinfo.txt")

	fmt.Println(string(data2))

}
