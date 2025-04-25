package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	variavel1 := 10
	variavel2 := &variavel1
	fmt.Println("variavel1:", variavel1)
	fmt.Println("variavel2:", variavel2)
	fmt.Println("variavel2:", *variavel2) // desreferenciação
	fmt.Println("-------")
	variavel1 = 20
	fmt.Println("variavel1:", variavel1)
	fmt.Println("variavel2:", variavel2)
	fmt.Println("variavel2:", *variavel2) // desreferenciação
	fmt.Println("-------")
}
