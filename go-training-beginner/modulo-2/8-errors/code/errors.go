package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := Divide(4, 0)

	if err != nil {
		fmt.Println(err) // Do something with the error
		return
	}

	fmt.Println(result) // Use the result
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	return a / b, nil
}
