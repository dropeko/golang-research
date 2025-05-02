package main

import "fmt"

type Carro struct {
	Marca  string
	Ano    int
	Modelo string
}

func (car Carro) ExibirDetalhes() {
	// O método ExibirDetalhes é um método associado ao tipo Carro.
	// Ele pode acessar os campos do tipo Carro e exibir seus detalhes.
	fmt.Printf("Marca: %s, Ano: %d, Modelo: %s\n", car.Marca, car.Ano, car.Modelo)
}

func (car *Carro) AtualizarAno(ano int) {
	// O método AtualizarAno é um método associado ao tipo Carro.
	// Ele pode modificar o campo Ano do tipo Carro.
	car.Ano = ano
}

func main() {
	fmt.Println("Métodos")
	fmt.Println("------------------------------")
	// Métodos são funções que estão associadas a um tipo específico.
	// Eles permitem que você defina comportamentos para tipos personalizados.
	// Isso é útil para encapsular dados e comportamentos relacionados em um único lugar.
	// Métodos são definidos com a palavra-chave func, seguida do nome do tipo ao qual o método pertence.
	// O método pode acessar os campos do tipo e modificar seu estado.
	// O método é chamado em uma instância do tipo, usando a sintaxe tipo.metodo().
	// Isso significa que o método pode acessar e modificar os campos do tipo.

	carro := Carro{
		Marca:  "Fusca",
		Ano:    1970,
		Modelo: "Fusca 1600",
	}

	carro.ExibirDetalhes()
	carro.AtualizarAno(2023)
	carro.ExibirDetalhes()
	fmt.Println("------------------------------")
}
