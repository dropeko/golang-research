package main

import "fmt"

func main() {
	fmt.Println("--------------")
	array1 := make([]int, 5, 10)
	fmt.Println(array1)

	array1 = append(array1, 77)
	fmt.Println(array1)
	fmt.Println("Capacidade do array:", cap(array1))
	fmt.Println("Tamanho do array:", len(array1))
	fmt.Println("--------------")

}
