package main

import (
	"fmt"
	enderecos "intro-testes/endereco"
)

func main(){
	tipoEndereco := enderecos.TipoDeEndereco("Avenida Paulista")
	fmt.Println(tipoEndereco)
}