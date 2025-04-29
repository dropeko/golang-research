package main

import "fmt"

func main() {
	fmt.Println("Funções Anônimas")
	// Função anônima
	soma := func(a, b int) int {
		return a + b
	}(5, 10) // Chama a função anônima com os argumentos 5 e 10

	fmt.Println("Resultado da soma:", soma)
}
