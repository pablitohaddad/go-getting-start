package main

import "fmt"


func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao	
	}
	return fibonacci(posicao - 2) + fibonacci(posicao - 1)
}

func main(){
	tarefas := make(chan int, 45)
	resultados := make(chan int, 45)

	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	for i := 0; i < 45; i++{
		tarefas <- i
	}
	close(tarefas)

	for i := 0; i < 45; i++{
		resultado := <- resultados
		fmt.Println(resultado)
	}

}

func worker(tarefas <- chan int, resultados chan <- int){ // Com as setas posso diferenciar quem envia e quem recebe dados. Tarefas recebe. Resultados envia
	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}
}