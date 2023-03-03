package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	u.nome = "Davi"
	u.idade = 21
	fmt.Println(u)

	usuario2 := usuario{"Davi", 21}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "davi"}
	fmt.Println(usuario3)
}
