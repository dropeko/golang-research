package main

import "fmt"

func main() {
	fmt.Println("Estruturas de Controle")

	numero := 11
	if numero > 10 {
		fmt.Println("O número é maior que 10")
	} else {
		fmt.Println("O número é menor ou igual a 10")
	}
	fmt.Println("-------")

	if outroNumero := 7; outroNumero > 10 { // <----- O outroNumero só existe dentro do escopo do if
		fmt.Println("O outro número é maior que 10")
	} else { // Pode usar else if
		fmt.Println("O outro número é menor ou igual a 10")
	}
	fmt.Println("-------")
}
