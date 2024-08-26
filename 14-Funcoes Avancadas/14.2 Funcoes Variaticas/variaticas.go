package main

import (
	"fmt"
)

func soma(numeros ...int) int{ // recebe de 1 a int numeros
	total := 0
	for _, numero := range numeros{
		total += numero
	}
	return total
}

func escrever (texto string, numeros ...int) {
	for _, numero := range numeros{
		fmt.Println(texto, numero)
	}
}

func main(){
	totalDaSoma := soma(1, 2, 3, 4)
	fmt.Println(totalDaSoma)

	escrever("Olá Mundo", 10, 20, 20, 30 , 20)


}