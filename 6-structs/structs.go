package main

import "fmt"

type user struct {
	nome     string
	idade    uint16
	endereco address
}

type address struct {
	logradouro string
	numero     uint16
}

func main() {
	fmt.Println("Structs")
	fmt.Println("-------")

	var u user
	u.nome = "Xablau"
	u.idade = 777
	fmt.Println(u)
	fmt.Println("-------")

	user1 := user{nome: "Dropeko", idade: 30}
	fmt.Println(user1)
	fmt.Println("-------")

	user2 := user{idade: 778}
	fmt.Println(user2)
	fmt.Println("-------")

	end := address{"Rua Paradise", 777}

	user3 := user{"Jolira", 18, end}
	fmt.Println(user3)
	fmt.Println("-------")

}
