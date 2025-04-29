package main

import "fmt"

func closure() func() {
	// A função closure retorna outra função que pode acessar as variáveis do escopo da função closure.
	// Isso significa que a função retornada pode acessar e modificar as variáveis definidas na função closure.
	// Isso é útil para criar funções que mantêm estado entre chamadas.
	// A variável texto é definida dentro da função closure e é acessada pela função retornada.
	texto := "Dentro da função closure"

	varFunc := func() {
		fmt.Println(texto)
	}

	return varFunc
}

func main() {
	// Funções closure são funções que capturam o ambiente em que foram criadas.
	// Elas podem acessar variáveis locais mesmo depois que a função que as criou já foi executada
	fmt.Println("Funções Closure")
	fmt.Println("------------------------------")

	texto := "Fora da função closure"
	fmt.Println(texto)

	retornoClosure := closure()
	retornoClosure()
	fmt.Println("------------------------------")
}
