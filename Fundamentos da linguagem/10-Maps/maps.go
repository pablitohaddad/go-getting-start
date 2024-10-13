package main

import "fmt"

func main(){
	fmt.Println("Maps")

	usuario := map[string]string { // [chave]valor e sempre tem que ser desse tipo
		"nome": "Pedro",
		"sobrenome": "Silva",
	}
	fmt.Println(usuario["nome"]) // acessa o valor da chave

	usuario2 := map[string]map[string]string{ // aninhamento de maps
		"nome":{
			"primeiro": "João",
			"ultimo": "Pedro",
		},
		"curso":{
			"nome": "Engenharia",
			"campus": "UNEMAT",
		},
	}
	fmt.Println(usuario2)
	delete(usuario2, "nome") // deleta a chave
	
	usuario2["signo"] = map[string]string{
		"nome": "Gêmeos",
	}
	fmt.Println(usuario2)
}