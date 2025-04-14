package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	soma := n1 + n2
	return soma
}

func somarMultiResult(n1 int8, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2

	return soma, subtracao
}

func main() {
	somaNums := somar(7, 7)
	fmt.Println(somaNums)

	var fVar = func(txt string) string {
		return txt
	}

	funcVar := fVar("Função por variavel")
	fmt.Println(funcVar)

	somaAd, somaSub := somarMultiResult(77, 17)
	fmt.Println(somaAd, somaSub)
}
