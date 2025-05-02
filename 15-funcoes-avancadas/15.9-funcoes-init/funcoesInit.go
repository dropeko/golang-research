package main

import "fmt"

func init() {
	fmt.Println("Executando função init")
	fmt.Println("------------------------------")
}

func main() {
	// Funções init são funções especiais em Go que são executadas automaticamente antes da função main.
	// Elas são úteis para inicializar pacotes ou executar código de configuração antes que o programa comece a rodar.

	// A função init é chamada automaticamente pelo Go antes da função main
	fmt.Println("Executando função main")
	fmt.Println("------------------------------")
}
