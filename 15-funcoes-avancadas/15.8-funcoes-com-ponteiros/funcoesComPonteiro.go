package main

import "fmt"

func modificaNumero(numero *int) {
	// A função modificaNumero recebe um ponteiro para um inteiro.
	// Isso significa que ela pode modificar o valor da variável original.
	// O operador * é usado para desreferenciar o ponteiro e acessar o valor armazenado no endereço de memória.
	*numero = *numero * -1
}

func main() {
	// Ponteiros são variáveis que armazenam o endereço de memória de outra variável.
	// Eles são úteis para passar grandes estruturas de dados para funções sem fazer cópias desnecessárias.
	// Além disso, eles permitem modificar o valor da variável original dentro da função.
	fmt.Println("Funções com Ponteiros")
	fmt.Println("------------------------------")

	numero := 10
	fmt.Println("Número antes da função:", numero)

	// Passando o endereço da variável numero para a função
	// Isso permite que a função modifique o valor original da variável
	modificaNumero(&numero)
	fmt.Println("Número depois da função:", numero)
	fmt.Println("------------------------------")
}
