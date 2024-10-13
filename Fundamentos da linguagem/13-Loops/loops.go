package main

import (
	"fmt"
	//"time"
)

func main(){
	i:=0

	for i < 10{
		i++
		fmt.Println("Incrementando i")
		// time.Sleep(time.Second) dormir por 1 segundo
	}
	fmt.Println(i)

	for j := 0; j < 10; j++{
		fmt.Println("Incrementando j", j)
		// time.Sleep(time.Second)
	}

	nomes := [3]string{"João", "Davi", "Lucas"}

	for indice, nome := range nomes { // for each
		fmt.Println(indice, nome)
	}

	for _, nome := range nomes { // para nao ter que colocar o indice
		fmt.Println(nome)
	}

	// nomeslice := []string{"João", "Davi", "Lucas"}

	for indice, letra := range "PALAVRA"{
		fmt.Println(indice, letra) // letra tem o valor da palavra na tabela ASCII
	}

	usuario := map[string]string{
		"nome":"Pablo",
		"sobrenome": "Haddad",
	}

	for chave, valor := range usuario{
		fmt.Println(chave, valor)
	}

	// Não podemos usar for em structs

	for{
		// for infinito
	}

}