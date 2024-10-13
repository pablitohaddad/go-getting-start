package main

import "fmt"


func recuperarExec(){
	if r := recover(); r != nil {
		fmt.Println("Exec recuperada")
	}
}

func alunoAprovado(n1, n2 float32) bool{
	defer recuperarExec()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6{
		return false
	}

	panic("A MÉDIOA É EXATAMENTE 6") // ele crasha o seu programa, fecha o seu programa, deve ser recuperada pelo recover. Diferente do error
	// retorna 0 pois o bool vira 0
}

func main(){

	fmt.Println(alunoAprovado(8, 6))
	fmt.Println("PÓS EXECUÇÃO")

}