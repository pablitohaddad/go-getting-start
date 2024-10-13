package main

import (
	"fmt"
	"testing"
)

func main(){

	soma := 1 + 2
	subtracao := 1 -2
	divisao := 10/4
	multiplicacao :=  10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)


	// var numero1 int16 = 10
	// var numero2 int32 = 25

	// somaDiferentes := numero1 + numero2 vai dar errado pois sao de tipos diferentes
	
	// FIM DOS ARITMETICOS

	var variavel1 string = "String"
	variavel2 := "String2"

	fmt.Println(variavel1, variavel2)

	// OPERADORES RELACIONAIS (true ou false)
 
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	// FIM DOS RELACIONAIS

	// OPERADORES LOGICOS
	fmt.Println("------------OPERADORES----------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!falso)
	fmt.Println(!verdadeiro)

	// OPERADORES UNÁRIOS

	numero := 10
	numero++
	fmt.Println(numero)
	numero += 15
	fmt.Println(numero)
	numero-= 15
	fmt.Println(numero)
	numero--
	fmt.Println(numero)

	numero *= 3
	numero /= 10
	numero %= 3

	// FIM dos operadores unarios

	// Operadores ternarios

	// text := numero > 5 ? "Maior que 5" : "Menor que 5" GO nao entende isso
	// premissa do do é de uma maneira só de fazer uma coisa
	// pq esse operadores ternario é algo que o IF faz, a unica coisa
	// que distoa é a declaração de variaveis

	var texto string
	if numero > 5{
		texto = "Maior que 5"
	}else{
		texto = "Menor que 5"
	}

	fmt.Println(texto)


}