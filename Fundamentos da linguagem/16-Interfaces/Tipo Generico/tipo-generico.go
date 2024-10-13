package main

import "fmt"

func generica(interf interface{}){ // qualquer coisa atende essa interface. com tipo generico
	fmt.Println(interf)
}

func main()  {

	generica("String")
	generica(123)
	generica(true)
	
	fmt.Println()

	mapa := map[interface{}]interface{}{
		1: "String",
		float32(100): false,
		true: float64(13),
	}

	fmt.Println(mapa)
}