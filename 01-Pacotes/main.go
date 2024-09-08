package main

import (
	"fmt"
	"modulo/auxiliar"
	"github.com/badoux/checkmail" // pacote externo
)


func main() {
	fmt.Println("Escrevendo do arquivo main")	
	auxiliar.Escrever() // minuscula = private maiuscula = public
	erro := checkmail.ValidateFormat("pablo@gmail.com")
	// go mod tidy para remover dependencias nao usadas
	fmt.Println(erro)
}