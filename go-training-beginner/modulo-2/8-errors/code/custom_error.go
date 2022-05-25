package main

import (
	"errors"
	"fmt"
)

type DivisionError struct {
	Code int
	Msg  string
}

func (d DivisionError) Error() string {
	return fmt.Sprintf("code %d: %s", d.Code, d.Msg)
}

func main() {
	result, err := Divide(4, 0)

	if err != nil {
		var divErr DivisionError

		switch {
		case errors.As(err, &divErr):
			fmt.Println(divErr) // Do something with the error
		default:
			fmt.Println("no idea!")
		}

		return
	}

	fmt.Println(result) // Use the result
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, DivisionError{
			Code: 2000,
			Msg:  "cannot divide by zero",
		}
	}

	return a / b, nil
}
