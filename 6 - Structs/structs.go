package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	var u usuario
	u.nome = "Bruno"
	u.idade = 20
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua dos bobos", 0}

	usuario2 := usuario{"Bruno", 20, enderecoExemplo}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Bruno"}
	fmt.Println(usuario3)
}
