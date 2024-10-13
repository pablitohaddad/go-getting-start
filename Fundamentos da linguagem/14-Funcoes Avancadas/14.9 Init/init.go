package main

import "fmt"

var n int

// Inicia antes da main, cada arquivo pode ter a sua init(). Usada para setup
func init(){
	fmt.Println("Função init sendo realizada")
	n = 10
}


func main(){
	fmt.Println("Função main")
	fmt.Println(n) // Fiz um setup dela na init
}