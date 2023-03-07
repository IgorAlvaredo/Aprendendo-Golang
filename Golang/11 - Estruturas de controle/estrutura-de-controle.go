package main

import "fmt"

func main() {
	fmt.Println("Estrutura de controle")

	numero := 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	if outronumero := numero; outronumero < 0 {
		fmt.Println("Numero maior que zero")
	} else if numero < -10 {
		fmt.Println("Numero menor que -10")
	} else {
		fmt.Println("Essa variavel esta ente 0 e -10")
	}

	fmt.Println()
}
