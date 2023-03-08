package main

import "fmt"

func main() {
	fmt.Println("Funções Anonimas")

	func() {
		fmt.Println("Olá Mundo!")
	}()

	func(texto string) {
		fmt.Println(texto)
	}("Passando Parametro")

	retorno := func(texto string) string {
		return fmt.Sprintf("Resebido -> %s", texto)
	}("Passando Parametro")

	fmt.Println(retorno)
}
