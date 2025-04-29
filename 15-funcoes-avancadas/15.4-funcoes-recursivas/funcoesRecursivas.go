package main

import "fmt"

func Fibonacci(n uint) uint {
	if n <= 1 {
		return n
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}

func main() {
	fmt.Println("Funções Recursivas")

	numero := uint(10)

	fmt.Println("-------")
	fmt.Printf("O %dº número de Fibonacci é: %d\n", numero, Fibonacci(numero))
	fmt.Println("-------")
}
