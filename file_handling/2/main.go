package main

import (
	"fmt"
	"os"
)

func main() {

	file, _ := os.Create("myinfo.txt")
	defer file.Close()

	fmt.Println("file created suceefully")

	data, _ := os.ReadFile("myinfo.txt")

	file.WriteString("my name is irfan")
	fmt.Println(string(data))
}
