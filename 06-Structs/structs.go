package main

import "fmt"

type usuario struct{
	nome string 
	idade uint8
	endereco endereco
}

type endereco struct{
	logradouro string
	numero uint8
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	fmt.Println(u)
	u.idade = 19
	u.nome = "Pablo"
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua dos bobos", 0}

	usuario2 := usuario{"Juan", 40, enderecoExemplo} // inferencia de tipo
	fmt.Println(usuario2)
	
	// usuario3 := usuario{"Pablo"} um erro q nn tem todos os valores

	usuario3 := usuario{nome: "Paulo"} // instanciando
	fmt.Println(usuario3)
}