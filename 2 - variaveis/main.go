package main

import "fmt"

func main() {
	// declaração explícita de variáveis
	var variavel1 string = "Variável 1"

	// declaração implícita de variáveis
	variavel2 := "Variável 2"

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	// declarando várias variáveis explicitamente
	var (
		variavel3 string = "Variável 3"
		variavel4 string = "Variável 4"
	)

	fmt.Println(variavel3, variavel4)

	// declarando várias variáveis implícitas
	variavel5, variavel6 := "Variável 5", "Variável 6"
	fmt.Println(variavel5, variavel6)

	// declararando constante
	const constante1 string = "Constante 1"
	fmt.Println(constante1)

	// troca de valores entre variáveis
	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)

}
