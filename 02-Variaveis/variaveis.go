package main

import "fmt"

func main(){
	var variavel1 string = "Variavel 1"
	variavel2 := "Variavel 2" // declaração implicita / inferencia de tipo
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "vairavel 3"
		variavel4 string = "variavel 4"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "Variavel 5", "Variavel 6"
	fmt.Println(variavel5, variavel6)

	// Varias maneiras de declarar uma variavel

	const const1 string = "sempre assim1"
	fmt.Println(const1)

	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)


}