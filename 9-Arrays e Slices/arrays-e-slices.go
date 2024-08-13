package main

import (
	"fmt"
	"reflect"
)

func main(){
	fmt.Println("Arrays e Slices")

	var array1[5] int // todos do mesmo tipo (5 posições)
	fmt.Println(array1) // Se vazio aplica valor 0 [0,0,0,0,0]
	array1[0] = 12
	array1[0] = 234
	array1[0] = 114
	array1[0] = 14
	array1[0] = 134

	array2 := [5]string{"Oi", "Ola", "Hello", "Hi", "Nice"} 
	fmt.Println(array2)

	array3 := [...] int{1, 2, 3, 4, 5} // calcula o numero q eu populei
	fmt.Println(array3)
	// array3[7] = 20 NÃO DA CERTO, ele é fixo

	// SLICE

	slice := []int{89234, 342134, 123424} // Declarando um slice, com tamanho dinamico, mas com limite de tipo
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice)) // retorna tipo da variavel
 	fmt.Println(reflect.TypeOf(array3))

	slice = append(slice, 123) //pega o slice, add o item novo e retorna outro slice com os valores atualizados
	fmt.Println(slice)

	slice2 := array2[0:3] // como um ponteiro que aponta para o iuntervalo de posições
	fmt.Println(slice2) // pega um pedaço do array e adiciona no slice2

	array2[1] = "Posição Alterada"
	fmt.Println(slice2) // se as posições mudarem, o slice também muda


	// Arrays Internos
	fmt.Println("===============================================================")
	slice3 := make([]float32, 10, 11) // cria array de 15 posicoes e retorna slice de 10 posicoes
	// tipo, tamanho e capacidade
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // length
	fmt.Println(cap(slice3)) // capacidade

	slice3 = append(slice3, 5) // se eu add mais 1 a capacidade muda pra 24
	// go cria outro array e dobra a capacidade dele, ex: se eu colocar o 12° valor, 
	// a capacidade vai pra 12 x 2 = 24

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	slice4 = append(slice4, 10)
	fmt.Println(len(slice4)) 
	fmt.Println(cap(slice4))
 }