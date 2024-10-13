package main

import "fmt"

func main(){

	fmt.Println("Estruturas de controle")

	numero := -5

	if numero > 15 {
		fmt.Println("Maior que 15")
	}else{
		fmt.Println("Menor igual a 15")
	}

	if outronumero := numero; outronumero > 0{ // if-init
		fmt.Println("Numero é maior que 0")
	}else if outronumero < - 10{
		fmt.Println("Numero é menor que 0")
	}else{
		fmt.Println("Entre - e -10")
	}

	// fmt.Println(outronumero) não da certo pq a variavel está somente no escopo do if


}