package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint16
}

type competidor struct {
	pessoa
	esporte    string
	campeao    bool
	aposentado bool
}

func main() {
	fmt.Println("Herança")

	pessoaContender := pessoa{"Dropeko", "Namek", 777}

	contender := competidor{pessoaContender, "CS2", true, true}
	fmt.Println(contender)
	fmt.Println("-------")
	fmt.Println(contender.nome)
	fmt.Println("-------")
}
