package main

import "fmt"

func recuperrarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso!")
	}
}

func alunoEstaAprovado(n1, n2 float64) bool {
	defer recuperrarExecucao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}
	panic("A Média é exatamente 6!")
}
func main() {
	fmt.Println("panic e Recover")
	fmt.Println(alunoEstaAprovado(8, 6))
	fmt.Println("Pós execução!")
}
