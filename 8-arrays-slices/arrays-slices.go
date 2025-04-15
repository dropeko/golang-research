package main

import "fmt"

func main() {
	fmt.Println("ARRAYS E SLICES")

	// ARRAYS
	var array1 [5]string
	array1[0] = "Dropeko"
	array1[1] = ""
	fmt.Println(array1)
	fmt.Println("-------")

	array2 := [5]string{"P1", "P2", "P3", "P4", "P5"}
	fmt.Println(array2)
	fmt.Println("-------")

	array3 := [...]int{7, 77, 777, 7777, 77777}
	fmt.Println(array3)
	fmt.Println("-------")

	// SLICES
	slice := []int{7, 17, 27, 37, 47, 57, 67, 77} // <----- Declaração básica do slice
	fmt.Println(slice)

	slice = append(slice, 87)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

}
