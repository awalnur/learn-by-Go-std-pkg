package main

import (
	"errors"
	"fmt"
)

var (
	ValidationError = errors.New("validation error")
	NotFoundError   = errors.New("not found error")
)

func getById(id int) (string, error) {
	if id == 1 {
		return "John", nil
	}
	return "", NotFoundError
}

func main() {

	name, err := getById(0)

	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Println("Validation error")
		} else if errors.Is(err, NotFoundError) {

			fmt.Println("Not found error")
		} else {
			fmt.Println("unknown error")
		}
	} else {
		fmt.Printf("Hello %s\n", name)
	}
}
