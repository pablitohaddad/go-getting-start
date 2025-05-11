package main

import (
	"fmt"
	"time"
)

// Concorrencia  != PARALELISMO



func main() {
	go escrever("Olá Mundo!") // goroutine 
	escrever("Programando em Go!")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
