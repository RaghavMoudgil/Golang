package main

import (
	"fmt"
)

type customError struct {
	Message string
	code    int
}

func (e *customError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.code, e.Message)
}

func main() {
	if 1 > 0 {
		err := &customError{
			Message: "An error occurred",
			code:    404,
		}
		fmt.Println(err.Error())
		return
	}
}
