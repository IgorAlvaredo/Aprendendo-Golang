package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int64 = 100000000000000
	fmt.Println(numero)

	var numero2 uint32 = 10000
	fmt.Println(numero2)

	//alias
	//INT32 = RUNE
	var numero3 rune = 123456
	fmt.Println(numero3)

	//BYTE = UINT8
	var numero4 rune = 123456
	fmt.Println(numero4)

	var numeroReal1 float32 = 1234.56
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 1234.56
	fmt.Println(numeroReal2)

	numeroReal3 := 1234.56
	fmt.Println(numeroReal3)

	//Fim numeros

	//String
	var string1 string = "Isso é uma string"
	fmt.Println(string1)

	string2 := "String 2"
	fmt.Println(string2)

	//retorna o numero referente a letra
	char := 'a'
	fmt.Println(char)

	//Fim string

	var bolleano bool = false
	fmt.Println(bolleano)

	var erro error = errors.New("Erro interno!")
	fmt.Println(erro)

}
