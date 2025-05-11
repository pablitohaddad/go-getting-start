package main

import (
	"fmt"
	"time"
)

func main(){

	canal := escrever("Olá mundo")

	for i := 0; i < 10; i++ {
		fmt.Println(<- canal)
	}


}

//  Encapsulando a chamada da ggo routine. Esse canal vai apenas esperar algo...
func escrever(texto string) <- chan string {
	canal := make(chan string)

	go func ()  {
		for{
			canal <- fmt.Sprintf("valor recebido %s", texto)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return canal
}