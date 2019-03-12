package utils

import (
	"fmt"
)

// Check Create panic if there is any error
func Check(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
