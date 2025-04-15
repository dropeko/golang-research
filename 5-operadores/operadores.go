package main

import "fmt"

func main() {
	// OPERADORES ARITMÉTICOS
	soma := 7 + 7
	subtracao := 7 - 7
	multiplicacao := 7 * 7
	divisao := 7 / 7
	modulo := 7 % 7

	fmt.Println(soma, subtracao, multiplicacao, divisao, modulo)

	var n1 int16 = 77
	var n2 int16 = 77
	var soma2 int16 = n1 + n2
	fmt.Println(soma2)
	fmt.Println("------")
	// FIM DOS ARITMÉTICOS

	// OPERADORES ATRIBUIÇÃO
	var atribuicao string = "Atribuição explícita"
	inferencia := "Inferência"

	fmt.Println(atribuicao, inferencia)
	fmt.Println("------")
	// FIM DOS OPERADORES ATRIBUIÇÃO

	// OPERADORES RELACIONAIS - SÓ RETORNAM TRUE OU FALSE
	verdade := true
	falso := false

	fmt.Println(verdade, falso)
	fmt.Println(10 > 7)
	fmt.Println(10 < 7)
	fmt.Println(10 >= 7)
	fmt.Println(10 <= 7)
	fmt.Println(10 == 7)
	fmt.Println(10 != 7)
	fmt.Println("------")
	// FIM DOS OPERADORES RELACIONAIS

	// OPERADORES LÓGICOS - TAMBÉM SÓ RETORNAM TRUE OU FALSE
	fmt.Println(verdade && falso)
	fmt.Println(verdade || falso)
	fmt.Println(!verdade)
	fmt.Println("------")
	// FIM DOS OPERADORES LÓGICOS

	// OPERADORES UNÁRIOS
	unarioNum := 777

	// Incremento e Decremento
	unarioNum++
	unarioNum--

	// Multi, Div, Mod
	unarioNum *= 10
	unarioNum /= 2
	unarioNum %= 5

	fmt.Println(unarioNum)
	fmt.Println("------")

}
