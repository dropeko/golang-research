package main

import "fmt"

func somaFunc(numeros ...int) int {
	soma := 0
	for _, numero := range numeros {
		soma += numero
	}
	return soma
}

func main() {
	fmt.Println("Funções Variáticas")

	fmt.Println(somaFunc(1, 2, 3, 4, 5))
	fmt.Println(somaFunc(1, 2, 3, 4, 5, 6, 7, 8, 9))
	fmt.Println(somaFunc(1))
	fmt.Println(somaFunc())
	fmt.Println("-------")

}
