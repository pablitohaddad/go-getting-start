package main

import "fmt"

type pessoa struct{
	nome string
	sobrenome string
	idade uint8
	altura uint8
}

type estudante struct{
	pessoa // herança em GO
	curso string
	faculdade string
}

func main(){
	fmt.Println("Heranca")

	p1 := pessoa{"Joao", "Pedro", 20, 178}
	fmt.Println(p1)

	e1 := estudante{p1, "Medicina", "USP"}
	fmt.Println(e1)


	fmt.Println(e1.nome, e1.sobrenome)

}