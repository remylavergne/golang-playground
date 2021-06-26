package main

import (
	"fmt"
	"testing"
)

func TestFunctionPointerBehavior(t *testing.T) {
	value := "Coucou"

	f := func(value *string) {
		fmt.Println(value)
	}

	for i := 0; i < 3; i++ {
		f(&value)
	}
}
