package main

import (
	"errors"
	"fmt"
)

func main() {
	// tipos de inteiros
	//int8, int16, int32, int64

	numero := -100000000000
	fmt.Println(numero)

	//uint == unassigned int
	var numero2 uint = 100000000000
	fmt.Println(numero2)

	//alias
	//rune = int32
	var numero3 rune = 123456
	fmt.Println(numero3)

	//byte = uint8
	var numero4 byte = 123
	fmt.Println(numero4)
	//fim números inteiros

	// início números reais
	//tipos de reais
	//float32, float64
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 1230000000.45
	fmt.Println(numeroReal2)

	numeroReal3 := 123.45
	fmt.Println(numeroReal3)

	// fim números reais

	// início strings
	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	char := 'A'
	fmt.Println(char)
	// fim strings

	// início booleanos
	var booleano1 bool = true
	fmt.Println(booleano1)
	// fim booleanos

	// início erro
	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
	// fim erro
}
