package main

import "fmt"

func main() {
	// DECLARAÇÃO BÁSICA DE MAPS
	maps := map[string]int{ // <----- Dentro do colchetes, o tipo da chave, fora do colchetes, o tipo do valor
		"idade": 1, // <----- Chave e valor
		"p2":    2, // <----- Chave e valor
		"p3":    3, // <----- Chave e valor
	}

	fmt.Println(maps)       // <----- Imprime o map
	fmt.Println(maps["p2"]) // <----- Imprime o valor da chave "p2"
	fmt.Println(maps["p4"]) // <----- Imprime o valor da chave "p4", que não existe, então imprime 0
	fmt.Println("-------")

	// Aninhando maps
	maps2 := map[string]map[string]string{ // <----- Dentro do colchetes, o tipo da chave, fora do colchetes, o tipo do valor
		"nome": {
			"primeiroNome": "Dropeko", // <----- Chave e valor
			"sobrenome":    "Namek",   // <----- Chave e valor
		},
	}
	fmt.Println(maps2)                      // <----- Imprime o map
	fmt.Println(maps2["nome"]["sobrenome"]) // <----- Imprime o valor da chave "idade" do map "p1"

	fmt.Println("-------")
	maps2["habilidade"] = map[string]string{ // <----- Adiciona um novo map dentro do map
		"mainHabilidade": "CS2", // <----- Chave e valor
	}
	fmt.Println(maps2)                                 // <----- Imprime o map
	fmt.Println(maps2["habilidade"]["mainHabilidade"]) // <----- Imprime o valor da chave "mainHabilidade" do map "habilidade"
}
