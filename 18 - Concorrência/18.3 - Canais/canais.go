package main

import (
	"fmt"
	"time"
)

//Canais podem enviar e receber dados, elas bloqueiam a exec do programa
func main(){
	canal := make(chan string)

	go escrever("Olá mundo", canal) 

	// Com for infinito da errado. Erro de DeadLock
	// for{
	// 	mensagem := <- canal // <- canal esperando receber um valor, ele fica esperando para poder passar essa linha. Esse cara espera.
	// 	fmt.Println(mensagem)
	// }	

	for{
		mensagem, aberto := <- canal // <- canal esperando receber um valor, ele fica esperando para poder passar essa linha. Esse cara espera.
		if !aberto {
			break
		}
	 	fmt.Println(mensagem)
	}	
	
	// Outra forma mais legível.
	for mensagem := range canal{
		fmt.Println(mensagem)
	}
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto // canal <- mandando um valor
		time.Sleep(time.Second)
	}

	close(canal)
}
