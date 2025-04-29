package main

import "fmt"

func alunoAprovado(nota float64) bool {
	// O defer é executado antes do retorno da funçao, na ordem inversa em que foi chamado.
	defer fmt.Println("Média calculada")
	fmt.Println("Verificando se o aluno está aprovado. Resultado:")
	if nota >= 7 {
		fmt.Println("Aprovado")
		return true
	}
	fmt.Println("Reprovado")
	return false
}

func main() {
	fmt.Println("Defer")
	fmt.Println("-------")

	// Defer é usado para adiar a execução de uma função até que a função que a contém retorne.
	// Isso é útil para garantir que certas ações sejam executadas, como liberar recursos ou fechar arquivos.
	// O defer é executado na ordem inversa em que foram chamados.
	resultado := alunoAprovado(9)
	fmt.Println("Resultado:", resultado)
}
