package main

import "fmt"

func main(){
	fmt.Println("Ponteiros")

	var var1 int = 10 
	var var2 int = var1 // copia

	fmt.Println(var1, var2)

	//PONTEIRO É UMA REFERENCIA DE MEMORIA

	var var3 int 
	var ponteiro *int // nomeclatura do ponteiro

	var3 = 100
	ponteiro = &var3 // endereço de memoria

	fmt.Println(var3, ponteiro) // lugar na memoria
	fmt.Println(var3, *ponteiro) // esreferenciação

	var3 = 150
	fmt.Println(var3, *ponteiro)
}