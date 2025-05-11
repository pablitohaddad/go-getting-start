package main

import "fmt"

func main(){
	canal := make(chan string, 2) // Capacidade 2, Buffer. Só bloqueia quando chegar a 2

	canal <- "Olá mundo !"
	canal <- "Programando em Go"
	// não printaria o Programando em Go
	// E se eu adicionar mais um, cai no DEADLOCK

	mensagem := <- canal
	fmt.Println(mensagem)
}