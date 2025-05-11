package main

import (
	"fmt"
	"sync"
	"time"
)


func main() {

	// Grupo de espera. Para sincronizar as funções
	var waitGroup sync.WaitGroup
	// Sempre especificar o numero de go rourines sendo usadas
	waitGroup.Add(2)

	go func(){
		escrever("Olá mundo")
		waitGroup.Done()
		// done da um -1 ao ADD
	}()

	go func(){
		escrever("Proigramando glang")
		waitGroup.Done()
	}()
	
	// espera chegar em 0
	waitGroup.Wait()
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
