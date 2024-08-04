package main

import "fmt"

func somar(n1 int8, n2 int8) int8 { // retorna o tipo dps do ()
	return n1 + n2
}

func calculoMatematicos(n1 int8, n2 int8) (int8, int8){
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main(){
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func (txt string) string{
		fmt.Println(txt)
		return txt
	}

	resultado := f("Texto da função")
	fmt.Println(resultado)

	resultadosSoma, resultadoSubtracao := calculoMatematicos(10, 15)
	fmt.Println(resultadosSoma, resultadoSubtracao)

	queroSoSoma, _ := calculoMatematicos(10, 15) // O _ para que ignore a subtração
	fmt.Println(queroSoSoma)

	_, queroSoSub:= calculoMatematicos(10, 15) // O _ para que ignore a soma
	fmt.Println(queroSoSub)

}