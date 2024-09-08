package main

import (
	"errors"
	"fmt"
)

func main() {
	//int8, int16, int32, int64, int (usa a arquitetura do seu pc)

	var numero int16 = 10000
	var numero2 int = -10000000 // vira 64x pois meu pc é de 64x bits
	fmt.Println(numero)
	fmt.Println(numero2)

	//Existe uint unsygned int (tem de 8 ate 64 tambem)

	//alias 
	var numero3 rune = 12345 // alias que representa o int32
	fmt.Println(numero3)
	//Para uint8, use byte como alias
	var numero4 byte = 123
	fmt.Println(numero4)


	// float32, float64

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 121231313423.45
	fmt.Println(numeroReal2)

	numeroReal3 := 12345.67
	fmt.Println(numeroReal3) //pega da sua arquitetura

	//String

	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	char := 'B' //entre aspas simples
	fmt.Println(char)

	// FIM STRING

	// Valor zero
	var texto int 
	fmt.Println(texto)

	var texto2 int 
	fmt.Println(texto2)

	//Boolean
	var booleano1 bool = true // false
	fmt.Println(booleano1) // valor zero = false

	//Error

	var erro error = errors.New("Erro interno") // pacote errors
	fmt.Println(erro)

	var erro2 error // valor zero = <nil>
	fmt.Println(erro2)
	

}