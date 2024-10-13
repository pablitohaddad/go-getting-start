package main

import "fmt"

func closure() func(){
	text := "Dentro da func closure"

	funcao := func()  {
		fmt.Println(text)
	}

	return funcao
}

func main()  {

	text := "Dentro da func main"
	fmt.Println(text)

	funcaoNova := closure()
	funcaoNova()

	
}