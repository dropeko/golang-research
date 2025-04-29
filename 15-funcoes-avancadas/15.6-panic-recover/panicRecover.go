package main

import "fmt"

func RecuperaFuncao() {
	if r := recover(); r != nil {
		println("Execução recuperada com sucesso")
	}
}

func AlunoAprovado(n1, n2 float32) bool {
	defer RecuperaFuncao()

	media := (n1 + n2) / 2

	if media > 6 {
		println("Aprovado")
		return true
	} else if media < 6 {
		fmt.Println("Recuperação")
	}

	panic("SIMULANDO UM PANIC!!!!!")
}

func main() {
	// O panic é uma forma de interromper a execução normal do programa e iniciar o processo de recuperação.
	// Ele é usado para indicar que ocorreu um erro grave que não pode ser tratado localmente.
	// O recover é uma função que permite recuperar de um panic e continuar a execução do programa.
	// Ele deve ser chamado dentro de uma função deferida para funcionar corretamente.

	// Exemplo de uso do panic e recover
	fmt.Println("Panic e Recover")
	fmt.Println("---------------------")

	resultado := AlunoAprovado(6, 6)
	fmt.Println("Resultado:", resultado)
	fmt.Println("---------------------")
}
