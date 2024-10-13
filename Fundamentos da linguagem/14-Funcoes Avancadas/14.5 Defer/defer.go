package main

import "fmt"


func f1(){
	fmt.Println("Exec 1")
}

func f2(){
	fmt.Println("Exec 2")
}

func alunoAprovado(n1, n2 float32) bool{
	defer fmt.Println("Média calculada") // antes do return, ele vai executar
	fmt.Println("Na funcao alunoAprovado")

	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}
	return false
}

func main(){
	defer f1() // deixa pra depois, vai ser a ultima a ser colocada e executada.
	f2()
	fmt.Println(alunoAprovado(7, 8))
}