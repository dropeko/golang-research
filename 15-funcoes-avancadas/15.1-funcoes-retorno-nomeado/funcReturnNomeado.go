package main

import "fmt"

func calculosMatematicos(n1, n2 int) (soma, subtracao, multiplicacao, divisao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	multiplicacao = n1 * n2
	divisao = n1 / n2
	return
}

func main() {
	// Retorno nomeado
	// O retorno nomeado é uma forma de retornar valores de uma função sem precisar usar a palavra-chave return
	// Isso é útil quando queremos retornar vários valores de uma função e não queremos usar a palavra-chave return
	fmt.Println("Retorno nomeado")
	fmt.Println("-------")
	soma, subtracao, multiplicacao, divisao := calculosMatematicos(10, 5)
	fmt.Println("Soma:", soma)
	fmt.Println("Subtração:", subtracao)
	fmt.Println("Multiplicação:", multiplicacao)
	fmt.Println("Divisão:", divisao)
	fmt.Println("-------")
}
