package main

import (
	"fmt"
)

func calculosMatematicos(n1, n2 int) (soma int, subtracao int){ // ja nomeio meu retorno
	soma = n1 + n2 // nao preciso do :=
	subtracao = n1 - n2
	return // apenas o return já é o suficiente
}

func main() {

	soma, subtracao := calculosMatematicos(10, 20)

	fmt.Println(soma, subtracao) // 30 -10

}