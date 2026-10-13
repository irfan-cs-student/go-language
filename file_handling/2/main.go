package main

import (
	"fmt"
	"os"
)

func main() {

	file, _ := os.Create("myinfo.txt")

	fmt.Println("file created suceefully")

	file.WriteString("my name is irfan")
	file.WriteString("\ni'm muslims")

	file.WriteString(`
i live in pakistan. 
I love creativity`)

	data, _ := os.ReadFile("myinfo.txt")

	fmt.Println(string(data))
	file.Close()

	err := os.Remove("myinfo.txt")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("file deleted")

}
