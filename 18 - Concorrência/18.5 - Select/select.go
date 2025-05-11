package main

import (
	"fmt"
	"time"
)

func main(){
	canal1, canal2 := make(chan string), make(chan string)

	go func ()  {
		for {
			time.Sleep(time.Millisecond * 500) // Sendo prejudicado. Tem que esperar o outro de 2 seg
			canal1 <- "Canal 1"	
		}
	}()

	go func ()  {
		for {
			time.Sleep(time.Millisecond * 2) // Delay maior
			canal2 <- "Canal 2"	
		}
	}()

	for {
		// Para igualar e "deixar justo"
		select{
		case msgCanal1 := <- canal1:
			fmt.Println(msgCanal1)
		case msgCanal2 := <- canal2:
			fmt.Println(msgCanal2)
		}
	}
}