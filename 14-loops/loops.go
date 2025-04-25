package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Loops")

	// for loop
	k := 0

	for k < 10 {
		fmt.Println("k:", k)
		k++
		time.Sleep(1 * time.Second)
	}
	fmt.Println("-------")

	for i := 0; i < 10; i += 2 {
		fmt.Println("i:", i)
	}
	fmt.Println("-------")

	// while loop
	i := 0
	for i < 10 {
		fmt.Println("i:", i)
		i++
	}
	fmt.Println("-------")

	// for range loop - Quando for iterar sobre um array, slice, string ou map
	numeros := []int{1, 2, 3, 4, 5}

	for index, numero := range numeros {
		fmt.Println("index:", index, "numero:", numero)
	}
	fmt.Println("-------")

	for index, letra := range "Dropeko" {
		fmt.Println("index:", index, "letra:", string(letra))
	}
	fmt.Println("-------")

	polePositions := map[string]string{
		"p1": "Dropeko",
		"p2": "Namek",
		"p3": "777",
	}

	for index, valor := range polePositions {
		fmt.Println("index:", index, "valor:", valor)
	}
	fmt.Println("-------")
}
