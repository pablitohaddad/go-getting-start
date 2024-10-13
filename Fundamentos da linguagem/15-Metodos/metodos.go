package main

import "fmt"

type usuario struct {
	nome string
	idade uint8
}

// sintaxe do método
func(u usuario) salvar(){
	fmt.Printf("Salvando os dados do usuario %s no banco de dados", u.nome)
}

func(u usuario) maiorIdade() bool{
	return u.idade >= 18
}

func (u *usuario) fazerAniversario(){
	u.idade++
}

func main()  {

	usuario1 := usuario{"Usuario 1", 20}
	
	usuario1.salvar()

	usuario2 := usuario{"davi", 30}
	usuario2.salvar()
	fmt.Println(usuario2.maiorIdade())

	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)
	
}