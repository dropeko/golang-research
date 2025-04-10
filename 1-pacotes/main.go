package main

import (
	"fmt"
	"pacotes-modulo/helper"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Writing from main ")
	helper.Write()
	helper.Write2()
	erro := checkmail.ValidateFormat("phcadev@gmail.com")
	fmt.Println(erro)
}
