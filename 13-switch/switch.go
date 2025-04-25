package main

import "fmt"

func diaDaSemana(dia int) string {
	switch dia { // Estrutura para avaliar o valor da variável dia
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-feira"
	case 3:
		return "Terça-feira"
	case 4:
		return "Quarta-feira"
	case 5:
		return "Quinta-feira"
	case 6:
		return "Sexta-feira"
	case 7:
		return "Sábado"
	default:
		return "Dia inválido"
	}
}

func diaDaSemana2(dia int) string {
	diaDaSemana := ""

	switch {
	case dia == 1:
		diaDaSemana = "Domingo"
	case dia == 2:
		diaDaSemana = "Segunda-feira"
	case dia == 3:
		diaDaSemana = "Terça-feira"
	case dia == 4:
		diaDaSemana = "Quarta-feira"
	case dia == 5:
		diaDaSemana = "Quinta-feira"
	case dia == 6:
		diaDaSemana = "Sexta-feira"
	case dia == 7:
		diaDaSemana = "Sábado"
	default:
		diaDaSemana = "Dia inválido"
	}

	return diaDaSemana
}

func main() {
	// O switch é uma estrutura de controle que permite executar um bloco de código com base no valor de uma variável ou expressão.
	// Ele é semelhante a uma série de if-else, mas é mais legível e fácil de entender.
	// O switch pode ser usado para comparar uma variável com vários valores diferentes, e o código correspondente ao valor correspondente
	fmt.Println("Switch")
	fmt.Println("Dia da semana:", diaDaSemana(1))
	fmt.Println("Dia da semana:", diaDaSemana(3))
	fmt.Println("Dia da semana:", diaDaSemana(5))
	fmt.Println("Dia da semana:", diaDaSemana(77))
	fmt.Println("--")
	fmt.Println("Dia da semana:", diaDaSemana2(1))
	fmt.Println("Dia da semana:", diaDaSemana2(3))
	fmt.Println("Dia da semana:", diaDaSemana2(5))
	fmt.Println("Dia da semana:", diaDaSemana2(77))
	fmt.Println("--")
}
