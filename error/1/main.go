package main

import (
	"errors"
	"fmt"
)

// error in go
func chekAge(a int) error {
	if a < 18 {
		return errors.New("age is less than 18")
	}
	return nil
}
func main() {

	err := chekAge(9)

	if err != nil {
		fmt.Println(err)
	}
}
