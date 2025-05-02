package main

import (
	"fmt"
	"math"
)

// Definindo uma interface
type forma interface {
	area() float64
}

// Definindo uma função que recebe uma interface como parâmetro
func calcularArea(f forma) {
	fmt.Printf("Calculando a área da forma: %0.2f\n", f.area())
}

// Definindo um tipo que implementa a interface
type retangulo struct {
	largura float64
	altura  float64
}

func (ret retangulo) area() float64 {
	// O método area calcula a área do retângulo
	return ret.largura * ret.altura
}

// Definindo outro tipo que implementa a interface
type circulo struct {
	radio float64
}

func (cir circulo) area() float64 {
	// O método area calcula a área do círculo
	return math.Pi * math.Pow(cir.radio, 2)
}

func main() {
	fmt.Println("Interfaces")
	fmt.Println("------------------------------")
	// Interfaces são tipos que definem um conjunto de métodos que um tipo deve implementar.
	// Elas permitem que você defina comportamentos comuns para diferentes tipos, sem precisar se preocupar com a implementação específica.
	// Isso é útil para criar código genérico e reutilizável.
	// Uma interface é definida com a palavra-chave type, seguida do nome da interface e da palavra-chave interface.
	// A interface define um conjunto de métodos que um tipo deve implementar.
	// Um tipo implementa uma interface se ele define todos os métodos da interface.
	// Isso significa que você pode usar qualquer tipo que implemente a interface como se fosse o tipo da interface.

	ret := retangulo{
		largura: 5,
		altura:  10,
	}
	calcularArea(ret)
	fmt.Println("------------------------------")

	cir := circulo{
		radio: 3,
	}
	calcularArea(cir)
	fmt.Println("------------------------------")

}
