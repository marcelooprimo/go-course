package main

import "fmt"

func main() {
	// ARITMETICOS
	soma := 10 + 20
	subtracao := 10 - 5
	divisao := 10 / 2
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)
	fmt.Println("-------------------------------")

	// FIM DOS ARITIMETICOS

	// ATRIBUICAO
	var variavel1 string = "String"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)
	fmt.Println("-------------------------------")
	// FIM DOS OPERADORES DE ATRIBUIÇÃO

	// OPERADORES RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)
	// FIM DOS OPERADORES RELACIONAIS

	// OPERADORES LÓGICOS
	fmt.Println("-------------------------------")
	fmt.Println(true && false) // false
	fmt.Println(true || false) // true
	fmt.Println(!true)         // false
	fmt.Println("-------------------------------")
	// FIM DOS OPERADORES LÓGICOS

	// OPERADORES UNÁRIOS
	numero := 10
	numero++     // numero = numero + 1
	numero += 15 // numero = numero + 15

	fmt.Println(numero)

	numero--     // numero = numero - 1
	numero -= 20 // numero = numero - 20

	numero *= 4 // numero = numero * 4
	numero /= 2 // numero = numero / 2
	numero %= 3 // numero = numero % 2
	fmt.Println(numero)
	fmt.Println("-------------------------------")
	// FIM DOS OPERADORES UNÁRIOS

	// OPERADORES TERNÁRIOS
	// EM GO É DIFERENTE DAS OUTRAS LINGUAGENS
	// OUTRAS LINGUAGENS -> variavel := condicao ? valor1 : valor2
	var texto string
	if numero > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}
	fmt.Println(texto)
	// FIM DOS OPERADORES TERNÁRIOS
}
