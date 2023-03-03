package main

import "fmt"

func main() {
	//Aritmeticos
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	var numero1 int16 = 10
	//var numero2 int32 = 25 nao pode
	var numero2 int16 = 25
	soma2 := numero1 + numero2

	fmt.Println(soma2)

	//atribuição
	var string string = "string1"
	strin2 := "string2"
	fmt.Println(string, strin2)

	//relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	//Logicos
	fmt.Println("-------------------------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	//Operadores unarios
	numero := 10
	numero++
	numero += 15
	fmt.Println(numero)

	numero--
	numero -= 20

	numero *= 3
	numero /= 10
	numero %= 3
	fmt.Println(numero)

	//operador ternario

	if numero > 5 {
		fmt.Println("numero maior que 5")
	} else {
		fmt.Println("numero menor que 5")
	}
}
