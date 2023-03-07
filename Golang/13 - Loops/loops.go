package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Loops")
	i := 0

	for i < 10 {
		i++
		fmt.Println("Incrementando 1")
		time.Sleep(time.Second)
	}

	fmt.Println(i)

	for j := 0; j < 10; j++ {
		fmt.Println("Incremento j", j)
		time.Sleep(time.Second)
	}

	nomes := [3]string{"João", "Davi", "Lucas"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	//Se não usar o indice
	for _, nome := range nomes {
		fmt.Println(nome)
	}

	for indice, letra := range "PALAVRA" {
		//Vai mostrar o codigo da letra
		fmt.Println(indice, letra)

		//Mostra a letra
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Leonardo",
		"Sobrenome": "Silva",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}
	//Com struct não vai

	for {
		fmt.Println("Execultando infinitamente")
		time.Sleep(time.Second)
	}
}
