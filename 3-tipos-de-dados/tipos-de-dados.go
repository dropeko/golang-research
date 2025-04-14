package main

import (
	"errors"
	"fmt"
)

func main() {
	// int8, int16, int 32 e int64
	numeroInteiro := 77
	fmt.Println(numeroInteiro)

	// float32 e float64
	numeroFloat := 77.7
	fmt.Println(numeroFloat)

	// string
	stringVar := "hello world"
	fmt.Println(stringVar)

	// boolean
	booleanVar := true
	fmt.Println(booleanVar)

	// error
	errorVar := errors.New("Novo erro")
	fmt.Println(errorVar)
}
